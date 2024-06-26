# jsondb

### Description

A simple json db in GO

### What it does?

This library enables to store, retrieve, update and delete data from db based on json files.

# Usage

#### func New

```go
func  New(path  string, divider  ...string) (*Jsondb, error)
```

New initializes a new json database at the specified path with an optional divider string.

#### type jsondb

```go
type  Jsondb  struct {
path     string
mux      *sync.RWMutex
divider  string
}
```

The jsondb type represents a JSON database.

## Insert

#### func (db \*jsondb) Insert

```go
func (db *jsondb) Insert(keysStr  string, value  any) error
```

Insert adds a new value to the JSON database using the specified keys.

#### func (db \*jsondb) InsertDir

```go
func (db *jsondb) InsertDir(dirsStr  string) error
```

InsertDir adds directories with the specified directory names if they do not exist.

## Get

#### func (db \*jsondb) GetLen

```go
func (db *Jsondb) GetLen(keysStr string) (int, error)
```
GetLen returns the number of entries in the directory.

#### func (db \*jsondb) GetAny

```go
func (db *jsondb) GetAny(keysStr  string) (any, error)
```

GetAny returns any stored under the specified key.

#### func (db \*jsondb) GetBool

```go
func (db *jsondb) GetBool(keysStr  string) (bool, error)
```

GetBool returns bool stored under the specified key, returns an error if the value cannot be converted to type bool.

#### func (db \*jsondb) GetInt

```go
func (db *jsondb) GetInt(keysStr  string) (int, error)
```

GetInt returns int stored under the specified key, returns an error if the value cannot be converted to type int.

#### func (db \*jsondb) GetFloat

```go
func (db *jsondb) GetFloat(keysStr  string) (float64, error)
```

GetFloat returns float64 stored under the specified key returns an error if the value cannot be converted to type float64.

#### func (db \*jsondb) GetString

```go
func (db *jsondb) GetString(keysStr  string) (string, error)
```

GetString returns string stored under the specified key, returns an error if the value cannot be converted to type string.

#### func (db \*jsondb) GetMap

```go
func (db *jsondb) GetMap(keysStr  string) (map[string]any, error)
```

GetMap returns map[string]any stored under the specified key, returns an error if the value cannot be converted to type map[string]any.

#### func (db \*jsondb) GetAllMaps

```go
func (db *jsondb) GetAllMaps(keysStr  string) ([]map[string]any, error)
```

GetAllMaps returns all files stored under the specified key converted to []map[string]any.

#### func (db \*jsondb) GetStruct

```go
func (db *jsondb) GetStruct(keysStr  string, dst  any) error
```

GetStruct returns a result stored under the specified key in the value pointed to by dst.

## Delete

#### func (db \*jsondb) Delete

```go
func (db *jsondb) Delete(keysStr  string) error
```

Delete deletes all values stored under the specified key.
