package jsondb

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// writeEntry writes an entry to the JSON database.
//
// Parameters:
// - entry: the entry to be written.
// - entryPath: the path where the entry will be written.
// Returns:
// - error: an error if any occurs during the write operation.
func (db *Jsondb) writeEntry(entry any, entryPath string) error {
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	err = os.WriteFile(entryPath, data, FileDefault)
	return err
}

// createDirs creates directories with the specified directory names if they do not exist.
//
// dirs: A slice of strings representing the names of the directories to be added.
// Returns the path of the last created directory and an error if any.
func (db *Jsondb) createDirs(dirs []string) (string, error) {
	path := db.path
	for _, dir := range dirs {
		path = filepath.Join(path, dir)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.Mkdir(path, 0755)
			if err != nil {
				return "", err
			}
		}
	}
	return path, nil
}
