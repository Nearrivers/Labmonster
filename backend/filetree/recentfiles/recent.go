package recentfiles

import (
	"bufio"
	"errors"
	"flow-poc/backend/config"
	"fmt"
	"os"
	"path/filepath"
	"slices"
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
		r.FilePaths = slices.DeleteFunc(r.FilePaths, func(p string) bool {
			return p == pathFromLabRoot
		})
	}

	// Cannot make use of r.FilePaths's capacity since this line resets it
	r.FilePaths = append([]string{pathFromLabRoot}, r.FilePaths...)
}

func (r *RecentlyOpened) LoadRecentlyOpended() error {
	p := r.getLabmonsterDirPath()
	fmt.Println(p)
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

func (r *RecentlyOpened) getLabmonsterDirPath() string {
	return filepath.Join(r.getLabPath(), ".labmonster", recentlyOpenedFilename)
}
