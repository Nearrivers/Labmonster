package dirhandler

import (
	"fmt"
	"os"
	"path/filepath"
)

func createNonDuplicateDir(absPath string) (string, error) {
	p := filepath.Dir(absPath)
	b := filepath.Base(absPath)

	for i := 1; ; i++ {
		name := fmt.Sprintf("%s %d", b, i)
		fp := filepath.Join(p, name)
		if doesDirExists(fp) {
			continue
		}

		err := os.Mkdir(fp, os.ModeAppend)
		if err != nil {
			return "", err
		}

		return name, nil
	}
}
