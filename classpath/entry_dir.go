package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (my *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(my.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, my, err
}

func (my *DirEntry) String() string {
	return my.absDir
}
