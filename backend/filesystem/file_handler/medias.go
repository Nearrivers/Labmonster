package file_handler

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

var (
	ErrMediaNotSupported  = errors.New("media type not supported")
	ErrCouldNotWriteMedia = errors.New("could not create media file")
)

// Given a path, absolute or relative to the lab's root, it will open a "media" file (images or videos)
// and will return the base64 encoded string to the client.
func (ft *FileHandler) OpenMedia(path string) (string, error) {
	var p string

	if filepath.IsAbs(path) {
		p = path
	} else {
		p = filepath.Join(ft.GetLabPath(), path)
	}

	b, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}

	m := mime.TypeByExtension(filepath.Ext(p))

	s := base64.StdEncoding.EncodeToString(b)
	return "data:" + m + ";base64," + s, nil
}

func (fh *FileHandler) SaveMedia(pathToFile, mimetype, base64File string) (string, error) {
	p := filepath.Dir(pathToFile)
	b, err := fileToBytes(base64File, mimetype)
	if err != nil {
		return "", err
	}

	fileName, err := fh.createFileName(p, mimetype)
	if err != nil {
		return "", err
	}

	f, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	n, err := f.Write(b)
	if err != nil {
		return "", err
	}

	if n == 0 {
		return "", ErrCouldNotWriteMedia
	}

	return f.Name(), nil
}

func fileToBytes(b64File, mimetype string) ([]byte, error) {
	s, _ := strings.CutPrefix(b64File, "data:"+mimetype+";base64,")
	b := make([]byte, base64.StdEncoding.DecodedLen(len(s)))
	n, err := base64.StdEncoding.Decode(b, []byte(s))
	if err != nil {
		return nil, err
	}

	if n == 0 {
		return nil, ErrNothingRead
	}

	return b, nil
}

func (fh *FileHandler) createFileName(pathFromLabRoot, mimetype string) (string, error) {
	t := time.Now()
	path := filepath.Join(fh.GetLabPath(), pathFromLabRoot, "Pasted Image "+t.Format("20060102150405"))

	switch mimetype {
	case "image/jpeg":
		return path + ".jpeg", nil
	case "image/png":
		return path + ".png", nil
	case "image/gif":
		return path + ".gif", nil
	case "image/webp":
		return path + ".webp", nil
	case "image/bmp":
		return path + ".bmp", nil
	case "video/mp4":
		return path + ".mp4", nil
	case "video/mpeg":
		return path + ".mpeg", nil
	case "video/webm":
		return path + ".webm", nil
	default:
		return "", ErrMediaNotSupported
	}
}

type MediaChunk struct {
	b      []byte
	offset int
}

const (
	mb    = 1024 * 1024
	limit = 20 * mb
)

func (fh *FileHandler) OpenMediaConc(path string) ([]byte, error) {
	log := logger.NewDefaultLogger()
	var p string

	if filepath.IsAbs(path) {
		p = path
	} else {
		p = filepath.Join(fh.GetLabPath(), path)
	}

	stat, err := os.Stat(p)
	if err != nil {
		return nil, err
	}

	fileSize := stat.Size()
	log.Error(fmt.Sprintf("taille du fichier en octets: %d", fileSize))
	b := make([]byte, 0)
	wg := sync.WaitGroup{}
	done := make(chan (bool), 1)
	fileContentCh := make(chan (MediaChunk))

	go func() {
		for fc := range fileContentCh {
			// b = slices.Insert(b, fc.offset, fc.b...)
			b = append(b, fc.b...)
		}

		done <- true
	}()

	var current int64
	var toReadLeft int64
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	for ; current < fileSize; current += limit + 1 {
		wg.Add(1)
		toReadLeft = fileSize - current

		go func(c, tr int64) {
			readMedia(f, c, tr, fileContentCh)
			log.Debug(fmt.Sprintf("Octets restant à lire: %d", tr))
			wg.Done()
		}(current, toReadLeft)
	}

	wg.Wait()
	close(fileContentCh)

	<-done
	close(done)
	return b, nil
}

func readMedia(f *os.File, offset, toReadLeft int64, ch chan<- (MediaChunk)) {
	f.Seek(offset, 0)
	reader := bufio.NewReader(f)
	var b []byte

	if toReadLeft > limit {
		b = make([]byte, limit)
	} else {
		b = make([]byte, toReadLeft)
	}

	_, readErr := reader.Read(b)
	if readErr != nil && readErr != io.EOF {
		panic(readErr)
	}

	ch <- MediaChunk{b, int(offset)}
}
