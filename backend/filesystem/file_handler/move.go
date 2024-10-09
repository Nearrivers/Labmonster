package file_handler

import (
	"io"
	"os"
)

// Move file utility function.
// If the new path is not already taken, the move operation will just move the file normally.
// Otherwise, it will create a non-duplicate file with a number at the end of its name. It will
// then copy the content of the old file into the new one to finally delete the old file.
// This function is taken as a file creation by the watcher.
func moveFile(oldPath, newPath string) (string, error) {
	if !doesFileExist(newPath) {
		err := os.Rename(oldPath, newPath)
		if err != nil {
			return "", nil
		}

		info, err := os.Stat(newPath)
		if err != nil {
			return "", err
		}
		return info.Name(), nil
	}

	oldFile, err := os.Open(oldPath)
	if err != nil {
		return "", err
	}

	newFile, name, err := createNonDuplicateFile(newPath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, oldFile)
	if err != nil {
		return "", err
	}

	err = oldFile.Close()
	if err != nil {
		return "", nil
	}

	// We delete the old file
	err = os.Remove(oldPath)
	if err != nil {
		return "", err
	}

	return name, err
}
