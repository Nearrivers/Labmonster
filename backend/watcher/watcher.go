package watcher

import (
	"context"
	"errors"
	"flow-poc/backend/config"
	"flow-poc/backend/filetree"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	// ErrDurationTooShort survient lorsque l'on appelle la méthode Start()
	// du watcher avec une durée inférieure à 1 nanoseconde
	ErrDurationTooShort = errors.New("error: duration is less than 1ns")

	// ErrWatcherRunning survient lorsque l'on tente d'appeller la méthode Start()
	// du watcher alors que le cycle de "polling" (sondage) est en train d'être exécuté
	// et que la méthode Close() n'a pas encore été appelée
	ErrWatcherRunning = errors.New("error: watcher is already running")

	// ErrWatcherFileDeleted survient lorsqu'un fichier / répertoire suivi est supprimé
	ErrWatcherFileDeleted = errors.New("error: watched file or folder deleted")

	// ErrSkip est une erreur utilisée afin de dire aux hooks des chemins de passer un fichier
	ErrSkip = errors.New("error: skipping file")
)

// Une Op est un type permettant de décrire quel type d'évènement
// est survenu lors d'un processus de surveillance
type Op uint32

// Opérations
const (
	Create Op = iota
	Write
	Remove
	Rename
	Chmod
	Move
)

var ops = map[Op]string{
	Create: "CREATE",
	Write:  "WRITE",
	Remove: "REMOVE",
	Rename: "RENAME",
	Chmod:  "CHMOD",
	Move:   "MOVE",
}

var FsOps = []struct {
	Value  Op
	TSName string
}{
	{Create, "CREATE"},
	{Remove, "REMOVE"},
	{Rename, "RENAME"},
	{Move, "MOVE"},
}

// Récupère une Op sous forme de string
func (e Op) String() string {
	if op, found := ops[e]; found {
		return op
	}
	return "???"
}

// Un Event décrit un évènement reçu lorsque des changements sur un fichier / répertoire surviennent.
// Cela inclut la os.FileInfo de l'élément modifié, le type d'évènement survenu ainsi que le chemin
// complet du fichier
type Event struct {
	Op       `json:"op"`
	Path     string            `json:"path"`
	OldPath  string            `json:"oldPath"`
	FilePath string            `json:"file"`
	FileType filetree.FileType `json:"fileType"`
	DataType filetree.DataType `json:"dataType"`
	os.FileInfo
}

// String retourne une string dépendant de quel type d'évènement est survenu ainsi que le fichier
// associé
func (e Event) String() string {
	if e.FileInfo == nil {
		return "???"
	}

	pathType := "FILE"
	if e.IsDir() {
		pathType = "DIRECTORY"
	}
	return fmt.Sprintf("%s %q %s %s %s [%s]", pathType, e.Name(), e.Op, e.FileType, e.DataType, e.Path)
}

// Marshal an event into something usable to the frontend
func (e *Event) MarshalFrontend(labpath string) {
	// Will error out if the old path is empty but it's fine
	op, err := filepath.Rel(labpath, e.OldPath)
	if err != nil {
		log.Printf("%s %s", err, e.OldPath)
	}

	p, err := filepath.Rel(labpath, e.Path)
	if err != nil {
		log.Printf("%s %s", err, e.Path)
	}

	e.OldPath = filepath.ToSlash(op)
	e.Path = filepath.ToSlash(filepath.Dir(p))
	e.FilePath = filepath.Base(p)
}

// Watcher décrit une process qui surveille les changements dans des fichiers
type Watcher struct {
	Event  chan Event
	Error  chan error
	Closed chan struct{}
	Ctx    context.Context
	close  chan struct{}
	wg     *sync.WaitGroup

	// mu protège les attributs le suivant
	config       *config.AppConfig
	mu           *sync.Mutex
	running      bool
	names        map[string]bool        // Booléan qui indique si on est récursif ou non
	files        map[string]os.FileInfo // map des fichiers
	ignored      map[string]struct{}    // map des fichiers / répertoires ignorés
	ignoreHidden bool                   // ignore les fichiers cachés
}

// New crée un nouveau Watcher
func New(cfg *config.AppConfig) *Watcher {
	// Setup du WaitGroup pour w.Wait
	var wg sync.WaitGroup
	wg.Add(1)

	return &Watcher{
		Event:   make(chan Event),
		Error:   make(chan error),
		Closed:  make(chan struct{}),
		close:   make(chan struct{}),
		config:  cfg,
		mu:      new(sync.Mutex),
		wg:      &wg,
		files:   make(map[string]os.FileInfo),
		ignored: make(map[string]struct{}),
		names:   make(map[string]bool),
	}
}

func (w *Watcher) SetContext(ctx context.Context) {
	w.Ctx = ctx
}

// AddRecursive ajoute un fichier ou un répertoire récursivement à la liste des fichiers
func (w *Watcher) AddRecursive(name string) (err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	absPath, err := filepath.Abs(name)
	if err != nil {
		return err
	}

	fileList, err := w.listRecursive(absPath)
	if err != nil {
		return err
	}

	for k, v := range fileList {
		w.files[k] = v
	}

	// Ajout du nom dans la liste des noms
	w.names[absPath] = true

	return nil
}

func (w *Watcher) listRecursive(name string) (map[string]os.FileInfo, error) {
	fileList := make(map[string]os.FileInfo)

	return fileList, filepath.Walk(name, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		_, ignored := w.ignored[path]

		isHidden, err := isHiddenFile(path)
		if err != nil {
			return err
		}

		// Si l'élément est ignoré ou que l'on ignore les éléments cachés
		if ignored || (w.ignoreHidden && isHidden) {
			// Si l'élément est un dossier
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Ajout du chemin et de l'os.FileInfo du fichier dans la liste des fichiers
		fileList[path] = info
		return nil
	})
}

// RemoveRecursive supprime soit un fichier soit un répertoire de manière récursive de la liste
// de fichiers
func (w *Watcher) RemoveRecursive(name string) (err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	name, err = filepath.Abs(name)
	if err != nil {
		return err
	}

	// supprime le nom de la liste des noms de w
	delete(w.names, name)

	// Si name est un simple fichier, le supprimer et retourner
	info, found := w.files[name]
	if !found {
		return nil // Le fichier n'existe pas, on peut retourner
	}

	if !info.IsDir() {
		delete(w.files, name)
		return nil
	}

	// Si c'est un répertoire, supprimer tout son contenu de w.files récursivement
	for path := range w.files {
		if strings.HasPrefix(path, name) {
			delete(w.files, path)
		}
	}

	return nil
}

// fileInfo est un mock de os.FileInfo utilisé pour déclencher des évènements manuellement
type fileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	sys     interface{}
	dir     bool
}

func (fs *fileInfo) IsDir() bool {
	return fs.dir
}
func (fs *fileInfo) ModTime() time.Time {
	return fs.modTime
}
func (fs *fileInfo) Mode() os.FileMode {
	return fs.mode
}
func (fs *fileInfo) Name() string {
	return fs.name
}
func (fs *fileInfo) Size() int64 {
	return fs.size
}
func (fs *fileInfo) Sys() interface{} {
	return fs.sys
}

func (w *Watcher) TriggerEvent(eventType Op, file os.FileInfo) {
	w.Wait()

	if file == nil {
		file = &fileInfo{name: "triggered event", modTime: time.Now()}
	}
	w.Event <- Event{Op: eventType, Path: "-", FileInfo: file}
}

func (w *Watcher) retrieveFileList() map[string]os.FileInfo {
	w.mu.Lock()
	defer w.mu.Unlock()

	fileList := make(map[string]os.FileInfo)

	for name := range w.names {
		list, err := w.listRecursive(name)
		if err != nil {
			if os.IsNotExist(err) {
				w.mu.Unlock()
				if name == err.(*os.PathError).Path {
					w.Error <- ErrWatcherFileDeleted
					w.RemoveRecursive(name)
				}
				w.mu.Lock()
			} else {
				// Erreur inconnue
				w.Error <- err
			}
		}

		for k, v := range list {
			fileList[k] = v
		}
	}

	return fileList
}

func (w *Watcher) Start(d time.Duration) error {
	// On retourne une erreur si le temps souhaité est inférieur à 1ns
	if d < time.Nanosecond {
		return ErrDurationTooShort
	}

	for {
		// On ne veut pas démarrer le watcher tant que le chemin vers le lab
		// n'est pas configuré
		if w.config.ConfigFile.LabPath != "" {
			if err := w.AddRecursive(w.config.ConfigFile.LabPath); err != nil {
				panic(err)
			}

			break
		}
	}

	// On vérifie si le Watcher tourne déjà
	w.mu.Lock()
	if w.running {
		w.mu.Unlock()
		return ErrWatcherRunning
	}
	w.running = true
	w.mu.Unlock()

	// Débloque w.Wait()
	w.wg.Done()

	for {
		// done signale au cycle de sondage interne lorsque la méthode du cycle
		// actuel a terminé son travail
		done := make(chan struct{})

		// Tout évènement trouvé est d'abord donné à evt avant d'être
		// envoyé channel Event principal
		evt := make(chan Event)

		// Récupère la liste des fichiers de tous les fichiers et répertoires surveillés
		fileList := w.retrieveFileList()

		// cancel peut être utilisé pour annuler la fonction de sondage actuelle
		cancel := make(chan struct{})

		// Recherche des évènements
		go func() {
			w.pollEvents(fileList, evt, cancel)
			done <- struct{}{}
		}()

	inner:
		for {
			select {
			case <-w.close:
				close(cancel)
				close(w.Closed)
				return nil
			case event := <-evt:
				w.Event <- event
			case <-done:
				break inner
			}
		}

		w.mu.Lock()
		w.files = fileList
		w.mu.Unlock()

		time.Sleep(d)
	}
}

func (w *Watcher) pollEvents(files map[string]os.FileInfo, evt chan Event, cancel chan struct{}) {
	w.mu.Lock()
	defer w.mu.Unlock()

	creates := make(map[string]os.FileInfo)
	removes := make(map[string]os.FileInfo)

	for path, info := range w.files {
		if _, found := files[path]; !found {
			removes[path] = info
		}
	}

	// Vérifie si un fichier a été créé, modifié et si un chmod est survenu
	for path, info := range files {
		_, found := w.files[path]
		if !found {
			// Un fichier a été créé
			creates[path] = info
		}
		// Je me fiche du reste je pense
	}

	// Boucle sur les 2 maps
	// Si un chemin est dans le map des opérations de suppressions ET dans le map
	// des opérations de créations alors cela veut dire que le fichier / répertoire
	// a été déplacé / renommé
	for path1, info1 := range removes {
		for path2, info2 := range creates {
			if !sameFile(info1, info2) || filepath.Dir(path1) != filepath.Dir(path2) {
				continue
			}

			var dType = filetree.FILE
			if info1.IsDir() {
				dType = filetree.DIR
			}

			// The move operation is now treated as a delete => create by default because of the non-duplicate creation.
			// When we move file x into directory y and a file with the same name already exists, the create file function
			// will create a brand new file with a different name, making the sameFile function just above say that the file
			// we just moved is different from the original.
			e := Event{
				Op:       Rename,
				Path:     path2,
				OldPath:  path1,
				FileInfo: info1,
				FileType: filetree.DetectFileType(filepath.Ext(path2)),
				DataType: dType,
			}

			// Si l'élément est toujours dans le même dossier, alors il a été renommé et non pas déplacé
			// if {
			// 	e.Op = Rename
			// }

			delete(removes, path1)
			delete(creates, path2)

			// Si le cycle n'est pas annulé, on écrit dans le channel evt
			select {
			case <-cancel:
				return
			default:
				evt <- e
			}
		}
	}

	// Envoie tous les autres évènements create et remove
	for path, info := range creates {
		select {
		case <-cancel:
			return
		default:
			var dType = filetree.FILE
			if info.IsDir() {
				dType = filetree.DIR
			}
			log.Println(path)
			e := Event{Create, path, "", "", filetree.DetectFileType(filepath.Ext(path)), dType, info}
			evt <- e
		}
	}

	for path, info := range removes {
		select {
		case <-cancel:
			return
		default:
			var dType = filetree.FILE
			if info.IsDir() {
				dType = filetree.DIR
			}
			log.Println(path)
			e := Event{Remove, path, path, "", filetree.DetectFileType(filepath.Ext(path)), dType, info}
			evt <- e
		}
	}
}

// Bloque jusqu'à que le watcher aie démarré
func (w *Watcher) Wait() {
	w.wg.Wait()
}

func (w *Watcher) Close() {
	w.mu.Lock()
	if !w.running {
		w.mu.Unlock()
		return
	}

	w.running = false
	w.files = make(map[string]os.FileInfo)
	w.names = make(map[string]bool)
	w.mu.Unlock()

	// Envoie un signal de close à la méthode Start()
	w.close <- struct{}{}
}
