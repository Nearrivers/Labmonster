package filetree

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	ErrMediaNotSupported  = errors.New("media type not supported")
	ErrCouldNotWriteMedia = errors.New("could not create media file")
)

func (ft *FileTree) SaveMedia(pathToFile, mimetype string, base64File string) (string, error) {
	p := filepath.Dir(pathToFile)
	b, err := fileToBytes(base64File, mimetype)
	if err != nil {
		return "", err
	}


	fileName, err := ft.createFileName(p, mimetype)
	if err != nil {
		return "", nil
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
	b, _ := strings.CutPrefix(b64File, "data:"+mimetype+";base64,")
	data := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	n, err := base64.StdEncoding.Decode(data, []byte(b))
	if err != nil {
		return nil, err
	}

	if n == 0 {
		return nil, ErrNothingRead
	}

	return data, nil
}

func (ft *FileTree) createFileName(pathFromLabRoot, mimetype string) (string, error) {
	t := time.Now()
	path := filepath.Join(ft.GetLabPath(), pathFromLabRoot,"Pasted Image "+ t.Format("20060102150405"))

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
	  return	path + ".mpeg", nil
	default:
		return "", ErrMediaNotSupported
	}
}
