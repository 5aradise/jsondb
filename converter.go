package jsondb

import "encoding/json"

// MapToStruct converts a map to a struct.
//
// dst: Pointer to destination variable to store the converted struct value.
// src: The map containing the data to be converted to a struct.
// error: Returns an error if any.
func MapToStruct(dst any, src map[string]any) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, dst)
	return err
}

// StructToMap converts a struct to a map.
//
// dst: The destination map variable to store the converted struct value.
// src: The struct containing the data to be converted to a map.
// error: Returns an error if any.
func StructToMap(dst map[string]any, src any) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &dst)
	return err
}
