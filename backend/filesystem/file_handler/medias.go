package file_handler

import (
	"encoding/base64"
	"errors"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"
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

func (fh *FileHandler) SaveMedia(fileName, pathToFile, mimetype, base64File string) (string, error) {
	p := filepath.Dir(pathToFile)
	b, err := fileToBytes(base64File, mimetype)
	if err != nil {
		return "", err
	}

	if fileName == "" {
		fn, err := fh.createFileName(p, mimetype)
		if err != nil {
			return "", err
		}

		fileName = fn
	} else {
		_, ext, err := getTypeAndExtensionWithMime(mimetype)
		if err != nil {
			return "", err
		}

		fileName = filepath.Join(fh.GetLabPath(), p, fileName) + ext
	}

	var f *os.File

	if doesFileExist(fileName) {
		f, _, err = createNonDuplicateFile(fileName)
	} else {
		f, err = os.Create(fileName)
	}

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
	var filename string
	t := time.Now().Format("20060102150405")

	mediaType, ext, err := getTypeAndExtensionWithMime(mimetype)
	if err != nil {
		return "", err
	}

	filename = mediaType + t + ext

	return filepath.Join(fh.GetLabPath(), pathFromLabRoot, filename), nil
}

func getTypeAndExtensionWithMime(mimetype string) (string, string, error) {
	var mediaType, extension string

	mediaType = "Image "
	switch mimetype {
	case "image/jpeg":
		extension = ".jpeg"
	case "image/png":
		extension = ".png"
	case "image/gif":
		extension = ".gif"
	case "image/webp":
		extension = ".webp"
	case "image/bmp":
		extension = ".bmp"
	case "video/mp4":
		mediaType = "Video "
		extension = ".mp4"
	case "video/mpeg":
		mediaType = "Video "
		extension = ".mpeg"
	case "video/webm":
		mediaType = "Video "
		extension = ".webm"
	default:
		return "", "", ErrMediaNotSupported
	}

	return mediaType, extension, nil
}
