package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}

	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}

func (my CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range my {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (my CompositeEntry) String() string {
	strs := make([]string, len(my))

	for i, entry := range my {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}
