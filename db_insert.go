package jsondb

import (
	"errors"
	"os"
	"path/filepath"
)

// Insert adds a new value to the JSON database using the specified keys.
//
// keysStr: A string representing the path where the value is stored.
// value: The value to be added.
// Returns an error if any.
func (db *Jsondb) Insert(keysStr string, value any) error {
	keys, err := parseArgs(keysStr)
	if err != nil {
		return err
	}

	db.mux.Lock()
	defer db.mux.Unlock()

	lastKey := keys[len(keys)-1]
	path := db.path
	if len(keys) >= 2 {
		dirPath, err := db.createDirs(keys[:len(keys)-2])
		if err != nil {
			return err
		}
		path = filepath.Join(dirPath, keys[len(keys)-2])
	}

	if _, err := os.Stat(path + DatabaseFileType); !os.IsNotExist(err) {
		entry, err := db.readEntry(path + DatabaseFileType)
		if err != nil {
			return errors.New("unreachable field")
		}

		entry[lastKey] = value
		return db.writeEntry(entry, path+DatabaseFileType)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
	}
	return db.writeEntry(value, filepath.Join(path, lastKey)+DatabaseFileType)

}

// InsertDir adds directories with the specified directory names if they do not exist.
//
// dirsStr: A string representing the names of the directories to be added.
// Returns an error if any.
func (db *Jsondb) InsertDir(dirsStr string) error {
	dirs, err := parseArgs(dirsStr)
	if err != nil {
		return err
	}

	db.mux.Lock()
	defer db.mux.Unlock()

	_, err = db.createDirs(dirs)
	return err
}
