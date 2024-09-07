package file_handler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Given an absolute path to a file we're trying to create. Before calling this function, we know that this file
// already exists. This function will try to create the same file but with a number appended to its name in order to avoid duplicates.
// It will continue to increment the appended number as long as the function finds a file with the same name. After succeeding in creating the file,
// it will return an os.File pointer to it that the caller will have to close, the actual name of the file to avoid f.Stat() boilerplate and an error
func createNonDuplicateFile(absPath string) (*os.File, string, error) {
	p := filepath.Dir(absPath)
	b := filepath.Base(absPath)
	newFileName := b[:strings.LastIndex(b, ".")]
	ext := filepath.Ext(absPath)

	for i := 1; ; i++ {
		name := fmt.Sprintf("%s %d%s", newFileName, i, ext)
		if doesFileExist(filepath.Join(p, name)) {
			continue
		}

		f, err := os.Create(filepath.Join(p, name))
		if err != nil {
			return nil, "", err
		}

		return f, name, nil
	}
}
