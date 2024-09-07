package file_handler

import "fmt"

type GetSubDirAndFilesError struct {
	err error
}

func (g *GetSubDirAndFilesError) Error() string {
	return fmt.Sprintf("couldn't build filetree: %v", g.err)
}

func (g *GetSubDirAndFilesError) Unwrap() error {
	return g.err
}

type WriteFileError struct {
	fileName string
	err      error
}

func (w *WriteFileError) Error() string {
	return fmt.Sprintf("could't write to file %s: %v", w.fileName, w.err)
}

func (w *WriteFileError) Unwrap() error {
	return w.err
}

type OpenFileError struct {
	err error
}

func (o *OpenFileError) Error() string {
	return fmt.Sprintf("couldn't open graph correctly: %v", o.err)
}

type SaveFileError struct {
	path string
	err  error
}

func (s *SaveFileError) Error() string {
	return fmt.Sprintf("couldn't save file %s: %v", s.path, s.err)
}

func (s *SaveFileError) Unwrap() error {
	return s.err
}
