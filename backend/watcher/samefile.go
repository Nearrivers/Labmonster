package watcher

import (
	"os"
)

func sameFile(fi1, fi2 os.FileInfo) bool {
	// return fi1.Name() == fi2.Name() &&
	return fi1.Size() == fi2.Size() &&
		fi1.Mode() == fi2.Mode() &&
		fi1.IsDir() == fi2.IsDir()
}
