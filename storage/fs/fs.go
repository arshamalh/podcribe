package fs

import (
	"io"
	"os"
	"path"
)

type FS struct {
	root string
}

func New(root string) *FS {
	return &FS{
		root: root,
	}
}

func (f *FS) Store(reader io.Reader, filename string, paths ...string) error {
	filePath := f.makeFullPath(filename, paths...)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, reader)
	return err
}

func (f *FS) Get(filename string, paths ...string) (io.ReadCloser, error) {
	filePath := f.makeFullPath(filename, paths...)
	return os.Open(filePath)
}

func (f *FS) makeFullPath(filename string, paths ...string) string {
	newPaths := []string{f.root}
	newPaths = append(newPaths, paths...)
	newPaths = append(newPaths, filename)
	return path.Join(newPaths...)
}
