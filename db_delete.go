package jsondb

import (
	"os"
	"path/filepath"
)


// Delete deletes all values stored under the specified key.
//
// keysStr: A string representing the path to the desired values.
// error: An error if any occurs during the delete operation.
// Returns: An error if any occurs during the delete operation.
func (db *Jsondb) Delete(keysStr string) error {
	keys, err := parseArgs(keysStr)
	if err != nil {
		return err
	}

	db.mux.Lock()
	defer db.mux.Unlock()

	pathKeys, lastKey := keys[:len(keys)-1], keys[len(keys)-1]
	path := filepath.Join(append([]string{db.path}, pathKeys...)...)

	entryPath := path + DatabaseFileType
	if _, err := os.Stat(entryPath); !os.IsNotExist(err) {
		entry, err := db.readEntry(entryPath)
		if err != nil {
			return err
		}

		delete(entry, lastKey)
		if len(entry) == 0 {
			return os.Remove(entryPath)
		}
		return db.writeEntry(entry, entryPath)
	}

	path = filepath.Join(path, lastKey)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Remove(path + DatabaseFileType)
	}

	return os.RemoveAll(path)
}
