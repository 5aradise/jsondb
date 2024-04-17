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

type jsondb struct {
	path    string
	mux     *sync.RWMutex
	divider string
}


// New initializes a new jsondb at the specified path with an optional divider string.
//
// path: the file path where the jsondb will be created.
// divider: an optional string used to divide nested keys in the jsondb.
// Returns a pointer to the newly initialized jsondb and an error.
func New(path string, divider ...string) (*jsondb, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, DirDefault)
		if err != nil {
			return nil, err
		}
	}

	db := jsondb{
		path:    path,
		mux:     &sync.RWMutex{},
		divider: ".",
	}

	if len(divider) != 0 {
		db.divider = divider[0]
	}

	return &db, nil
}
