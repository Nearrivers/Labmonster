package filetree

import (
	"encoding/base64"
	"errors"
	"os"
	"time"
)

var (
	ErrMediaNotSupported  = errors.New("media type not supported")
	ErrCouldNotWriteMedia = errors.New("could not create media file")
)

func (ft *FileTree) SaveMedia(pathFromLabRoot, mimetype string, b string) (string, error) {
	// b, err := fileToBytes(base64File)
	// if err != nil {
	// 	return "", err
	// }

	var f *os.File
	var err error
	defer f.Close()

	fileName := createFileName()

	switch mimetype {
	case "image/jpeg":
		f, err = os.Create(fileName + ".jpeg")
	case "image/png":
		f, err = os.Create(fileName + ".png")
	default:
		return "", ErrMediaNotSupported
	}

	if err != nil {
		return "", err
	}

	n, err := f.Write(b)
	if err != nil {
		return "", err
	}

	if n == 0 {
		return "", ErrCouldNotWriteMedia
	}

	return f.Name(), nil
}

func fileToBytes(b64File string) ([]byte, error) {
	data := make([]byte, base64.StdEncoding.DecodedLen(len(b64File)))
	n, err := base64.StdEncoding.Decode(data, []byte(b64File))
	if err != nil {
		return nil, err
	}

	if n == 0 {
		return nil, ErrNothingRead
	}

	return data, nil
}

func createFileName() string {
	t := time.Now()
	return t.Format("260020060102150405")
}
