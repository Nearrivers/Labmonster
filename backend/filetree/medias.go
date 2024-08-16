package filetree

import (
	"bufio"
	"encoding/base64"
	"errors"
	"io"
	"mime"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"
)

type MediaChunk struct {
	b      []byte
	offset int
}

var (
	ErrMediaNotSupported  = errors.New("media type not supported")
	ErrCouldNotWriteMedia = errors.New("could not create media file")
)

// Given a path, absolute or relative to the lab's root, it will open a "media" file (images or videos)
// and will return the base64 encoded string to the client.
func (ft *FileTree) OpenMedia(path string) (string, error) {
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

func (ft *FileTree) SaveMedia(pathToFile, mimetype string, base64File string) (string, error) {
	p := filepath.Dir(pathToFile)
	b, err := fileToBytes(base64File, mimetype)
	if err != nil {
		return "", err
	}

	fileName, err := ft.createFileName(p, mimetype)
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

func (ft *FileTree) createFileName(pathFromLabRoot, mimetype string) (string, error) {
	t := time.Now()
	path := filepath.Join(ft.GetLabPath(), pathFromLabRoot, "Pasted Image "+t.Format("20060102150405"))

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
	default:
		return "", ErrMediaNotSupported
	}
}

const (
	mb    = 1024 * 1024
	limit = 20 * mb
)

func (ft *FileTree) OpenMediaConc(path string) (string, error) {
	var p string

	if filepath.IsAbs(path) {
		p = path
	} else {
		p = filepath.Join(ft.GetLabPath(), path)
	}

	wg := sync.WaitGroup{}

	f, err := os.Open(p)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(f)

	return "", nil
}

// func (ft *FileTree) OpenMediaConc(path string) (string, error) {
// 	var p string

// 	if filepath.IsAbs(path) {
// 		p = path
// 	} else {
// 		p = filepath.Join(ft.GetLabPath(), path)
// 	}

// 	stat, err := os.Stat(p)
// 	if err != nil {
// 		return "", err
// 	}

// 	fileSize := stat.Size()
// 	b := make([]byte, fileSize)
// 	wg := sync.WaitGroup{}
// 	done := make(chan (bool), 1)
// 	fileContentCh := make(chan (MediaChunk))

// 	go func() {
// 		for fc := range fileContentCh {
// 			b = slices.Insert(b, fc.offset, fc.b...)
// 		}

// 		done <- true
// 	}()

// 	var current int64
// 	f, err := os.Open(p)
// 	if err != nil {
// 		return "", err
// 	}

// 	defer f.Close()
// 	for ; current < fileSize; current += limit + 1 {
// 		wg.Add(1)

// 		go func(c int64) {
// 			readMedia(f, c, fileContentCh)
// 			wg.Done()
// 		}(current)
// 	}

// 	wg.Wait()
// 	close(fileContentCh)

// 	<-done
// 	close(done)

// 	m := mime.TypeByExtension(filepath.Ext(p))

// 	s := base64.StdEncoding.EncodeToString(b)
// 	return "data:" + m + ";base64," + s, nil
// }

func readMedia(f *os.File, offset int64, ch chan (MediaChunk)) {
	f.Seek(offset, 0)
	reader := bufio.NewReader(f)
	b := make([]byte, limit)

	_, readErr := reader.Read(b)
	if readErr != nil && readErr != io.EOF {
		panic(readErr)
	}

	ch <- MediaChunk{b, int(offset)}
}
