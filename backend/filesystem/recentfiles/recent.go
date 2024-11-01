package recentfiles

import (
	"bufio"
	"errors"
	"flow-poc/backend/config"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const (
	recentlyOpenedFilename = "recentlyOpened.txt"
)

// struct that handle recently opened files
// The application will remember the last n files opened
// where n is equal to the maxFile key just above
type RecentlyOpened struct {
	Cfg       *config.AppConfig
	FilePaths []string
	maxFiles  int
}

func NewRecentlyOpened(c *config.AppConfig, max int) *RecentlyOpened {
	return &RecentlyOpened{c, make([]string, 0), max}
}

func (r *RecentlyOpened) getLabPath() string {
	return r.Cfg.ConfigFile.LabPath
}

func (r *RecentlyOpened) GetRecentlyOpenedFiles() ([]string, error) {
	if len(r.FilePaths) == 0 {
		err := r.LoadRecentlyOpended()
		if err != nil {
			return nil, err
		}
	}
	return r.FilePaths, nil
}

// Truncate recently opended text file and save the paths contained
// in the FilePaths array
func (r *RecentlyOpened) SaveRecentlyOpended() error {
	p := r.getLabmonsterDirPath()
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, p := range r.FilePaths {
		_, saveErr := fmt.Fprintln(f, p)
		if saveErr != nil {
			return err
		}
	}

	return nil
}

// Prepend a path relative to the lab's root to the FilePath array. A path will not be present
// twice in this array. RecentlyOpened.maxFiles sets the maximum of recently opened files
// and this function will make sure the capacity is never exceeded.
func (r *RecentlyOpened) AddRecentFile(pathFromLabRoot string) {
	if len(r.FilePaths) == r.maxFiles {
		r.FilePaths = r.FilePaths[0 : len(r.FilePaths)-1]
	}

	if slices.Contains(r.FilePaths, pathFromLabRoot) {
		r.RemoveRecent(pathFromLabRoot)
	}

	// Cannot make use of r.FilePaths's capacity since this line resets it
	r.FilePaths = append([]string{pathFromLabRoot}, r.FilePaths...)
}

// Replace a recent file with a new one. This method is used when renaming a file to make sure
// the file can still be opened via the recent file command
func (r *RecentlyOpened) ReplaceRecent(oldPath, newPath string) {
	i := slices.Index(r.FilePaths, oldPath)
	if i == -1 {
		return
	}

	r.FilePaths = slices.Replace(r.FilePaths, i, i, newPath)
}

// Remove a recent file. Used when deleting a file
func (r *RecentlyOpened) RemoveRecent(pathFromLabRoot string) {
	r.FilePaths = slices.DeleteFunc(r.FilePaths, func(p string) bool {
		return p == pathFromLabRoot
	})
}

// Read the recentFile.txt and load each line inside the RecentlyOpened.FilePaths array
func (r *RecentlyOpened) LoadRecentlyOpended() error {
	p := r.getLabmonsterDirPath()
	f, err := os.Open(p)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		r.FilePaths = append(r.FilePaths, scanner.Text())
	}

	if errScan := scanner.Err(); errScan != nil {
		return errScan
	}

	return nil
}

// This function will try to open every recently opened file. If the file doesn't exists
// anymore, it will remove it from the list. This function is called after deleting a directory
// in order to check if any recent files were in that directory.
func (r *RecentlyOpened) CheckIfRecentFileStillExists() {
	labPath := r.getLabPath()

	for _, recentFile := range r.FilePaths {
		path := filepath.Join(labPath, recentFile)

		f, err := os.Open(path)
		if err != nil && os.IsNotExist(err) {
			r.RemoveRecent(recentFile)
		}

		f.Close()
	}
}

// Iterates over recently opened file list and replace every instance
// of oldPathFromRoot with newPathFromRoot. This function is used after renaming a directory
// to make sure every paths are still relevant with the user's machine.
func (r *RecentlyOpened) ReconcilePaths(oldPathFromRoot, newPathFromRoot string) {
	for i, recentFile := range r.FilePaths {
		if !strings.Contains(recentFile, oldPathFromRoot) {
			continue
		}

		line := strings.Replace(recentFile, oldPathFromRoot, newPathFromRoot, 1)
		r.FilePaths[i] = line
	}
}

func (r *RecentlyOpened) getLabmonsterDirPath() string {
	return filepath.Join(r.getLabPath(), ".labmonster", recentlyOpenedFilename)
}
