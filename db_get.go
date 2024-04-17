package jsondb

import (
	"os"
	"path"
)

// GetLen retrieves the number of entries in a directory specified by the given keys.
//
// keysStr: A string representing the path to the directory.
// Returns the number of entries in the directory.
func (db *Jsondb) GetLen(keysStr string) (int, error) {
	keys, err := parseArgs(keysStr)
	if err != nil {
		return 0, err
	}

	db.mux.RLock()
	defer db.mux.RUnlock()

	f, err := os.Open(path.Join(append([]string{db.path}, keys...)...))
	if err != nil {
		return 0, err
	}
	files, err := f.Readdir(0)
	if err != nil {
		return 0, err
	}

	return len(files), nil
}

// GetAny retrieves an any value from the JSON database based on the specified keys.
//
// keysStr: A string representing the path to the desired integer value.
// Returns the any value found at the given keys and an error if not.
func (db *Jsondb) GetAny(keysStr string) (any, error) {
	keys, err := parseArgs(keysStr)
	if err != nil {
		return nil, err
	}

	db.mux.RLock()
	defer db.mux.RUnlock()

	return db.read(keys)
}

// GetInt retrieves an integer value from the JSON database based on the specified keys.
//
// keysStr: A string representing the path to the desired integer value.
// Returns the integer value found at the given keys and an error if any.
func (db *Jsondb) GetInt(keysStr string) (int, error) {
	keys, err := parseArgs(keysStr)
	if err != nil {
		return 0, err
	}

	db.mux.RLock()
	defer db.mux.RUnlock()

	val, err := db.read(keys)
	if err != nil {
		return 0, err
	}
	if val, ok := val.(float64); ok && val == float64(int(val)) {
		return int(val), nil
	}
	return 0, ErrorTypeConverting
}

// GetFloat retrieves an float value from the JSON database based on the specified keys.
//
// keysStr: A string representing the path to the desired integer value.
// Returns the float64 value found at the given keys and an error if any.
func (db *Jsondb) GetFloat(keysStr string) (float64, error) {
	keys, err := parseArgs(keysStr)
	if err != nil {
		return 0, err
	}

	db.mux.RLock()
	defer db.mux.RUnlock()

	val, err := db.read(keys)
	if err != nil {
		return 0, err
	}
	if val, ok := val.(float64); ok {
		return val, nil
	}
	return 0, ErrorTypeConverting
}

// GetString retrieves an integer value from the JSON database based on the specified keys.
//
// keysStr: A string representing the path to the desired integer value.
// Returns the string value found at the given keys and an error if any.
func (db *Jsondb) GetString(keysStr string) (string, error) {
	keys, err := parseArgs(keysStr)
	if err != nil {
		return "", err
	}

	db.mux.RLock()
	defer db.mux.RUnlock()

	val, err := db.read(keys)
	if err != nil {
		return "", err
	}
	if val, ok := val.(string); ok {
		return val, nil
	}
	return "", ErrorTypeConverting
}

// GetBool retrieves a boolean value from the JSON database based on the specified keys.
//
// keysStr: A string representing the path to the desired integer value.
// Returns the boolean value found at the given keys and an error if any.
func (db *Jsondb) GetBool(keysStr string) (bool, error) {
	keys, err := parseArgs(keysStr)
	if err != nil {
		return false, err
	}

	db.mux.RLock()
	defer db.mux.RUnlock()

	val, err := db.read(keys)
	if err != nil {
		return false, err
	}
	if val, ok := val.(bool); ok {
		return val, nil
	}
	return false, ErrorTypeConverting
}

// GetMap retrieves a map value from the JSON database based on the specified keys.
//
// keysStr: A string representing the path to the desired integer value.
// Returns the map value found at the given keys and an error if any.
func (db *Jsondb) GetMap(keysStr string) (map[string]any, error) {
	keys, err := parseArgs(keysStr)
	if err != nil {
		return nil, err
	}

	db.mux.RLock()
	defer db.mux.RUnlock()

	val, err := db.read(keys)
	if err != nil {
		return nil, err
	}
	if val, ok := val.(map[string]any); ok {
		return val, nil
	}
	return nil, ErrorTypeConverting
}

// GetStruct retrieves a struct value from the JSON database based on the specified keys.
//
// keysStr: A string representing the path to the desired integer value.
// dst: Pointer to destination variable to store the retrieved struct value.
// Returns an error if any.
func (db *Jsondb) GetStruct(keysStr string, dst any) error {
	val, err := db.GetMap(keysStr)
	if err != nil {
		return err
	}

	return MapToStruct(dst, val)
}

// GetAllMaps retrieves all map values from the JSON database based on the specified keys.
//
// keysStr: A string representing the path to the desired integer value.
// Returns an array of map values found at the given keys and an error if any.
func (db *Jsondb) GetAllMaps(keysStr string) ([]map[string]any, error) {
	keys, err := parseArgs(keysStr)
	if err != nil {
		return nil, err
	}

	db.mux.RLock()
	defer db.mux.RUnlock()

	val, err := db.readAll(keys)
	if err != nil {
		return nil, err
	}
	if val, ok := val.([]map[string]any); ok {
		return val, nil
	}
	return nil, ErrorTypeConverting
}
