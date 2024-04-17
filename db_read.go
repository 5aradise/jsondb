package jsondb

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// read reads the value at the specified keys in the JSON database.
//
// keys: A slice of strings representing the path to the desired value.
// Returns any type of value found at the given keys and an error if any.
func (db *Jsondb) read(keys []string) (any, error) {
	pathKeys, lastKey := keys[:len(keys)-1], keys[len(keys)-1]

	path := filepath.Join(append([]string{db.path}, pathKeys...)...)

	if _, err := os.Stat(path); err == nil {
		path = filepath.Join(path, lastKey)

		entry, err := db.readEntry(path + DatabaseFileType)
		if err != nil {
			return nil, err
		}

		return entry, nil
	}

	entry, err := db.readEntry(path + DatabaseFileType)
	if err != nil {
		return nil, errors.New("entry not found")
	}

	value, ok := entry[lastKey]
	if !ok {
		return nil, errors.New("value not found")
	}

	return value, nil
}

// readEntry reads the entry at the specified entryPath in the JSON database.
//
// entryPath: A string representing the path to the desired entry.
// Returns the entry found at the entryPath and an error if any.
func (db *Jsondb) readEntry(entryPath string) (entry map[string]any, err error) {
	data, err := os.ReadFile(entryPath)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &entry)
	return
}

// readAll reads all entries for the specified keys in the JSON database.
//
// keys: A slice of strings representing the path to the desired entries.
// Returns a slice of entries found at the given keys and an error if any.
func (db *Jsondb) readAll(keys []string) (any, error) {
	path := filepath.Join(append([]string{db.path}, keys...)...)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	files, err := f.Readdir(0)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	entrys := make([]map[string]any, 0, len(files))

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		entry, err := db.readEntry(filepath.Join(path, file.Name()))
		if err != nil {
			return nil, err
		}
		entrys = append(entrys, entry)
	}
	return entrys, nil
}
