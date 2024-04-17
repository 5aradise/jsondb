package jsondb

// GetAny retrieves an any value from the JSON database based on the specified keys.
//
// keysStr: A string representing the path to the desired integer value.
// Returns the any value found at the given keys and an error if not.
func (db *jsondb) GetAny(keysStr string) (any, error) {
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
func (db *jsondb) GetInt(keysStr string) (int, error) {
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
func (db *jsondb) GetFloat(keysStr string) (float64, error) {
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
func (db *jsondb) GetString(keysStr string) (string, error) {
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
func (db *jsondb) GetBool(keysStr string) (bool, error) {
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
func (db *jsondb) GetMap(keysStr string) (map[string]any, error) {
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
// dst: The destination variable to store the retrieved struct value.
// Returns an error if any.
func (db *jsondb) GetStruct(keysStr string, dst any) error {
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
func (db *jsondb) GetAllMaps(keysStr string) ([]map[string]any, error) {
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
