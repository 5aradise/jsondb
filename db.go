package jsondb

import (
	"errors"
	"os"
	"sync"
)

const (
	DatabaseFileType = ".json"
	DirDefault       = 0755
	FileDefault      = 0644
)

var (
	ErrorTypeConverting = errors.New("can't convert type")
)

type Jsondb struct {
	path    string
	mux     *sync.RWMutex
	divider string
}

// New initializes a new json database at the specified path with an optional divider string.
//
// path: the file path where the jsondb will be created.
// divider: an optional string used to divide nested keys in the jsondb.
// Returns a pointer to the newly initialized Jsondb instance and an error.
func New(path string, divider ...string) (*Jsondb, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, DirDefault)
		if err != nil {
			return nil, err
		}
	}

	db := Jsondb{
		path:    path,
		mux:     &sync.RWMutex{},
		divider: ".",
	}

	if len(divider) != 0 {
		db.divider = divider[0]
	}

	return &db, nil
}

func (db *Jsondb) Divider() string {
	return db.divider
}
