package levelsets

import (
	"reflect"
)

// LevelSet a struct to hold a levelset.
type LevelSet struct {
	Name    string
	Credit  string
	License string
	Levels  []Level
}

// Level - a struct to hold a level.
type Level struct {
	Name     string
	ID       string
	Formula  string
	Atoms    map[byte]([]byte)
	Arena    [][]byte
	Molecule [][]byte
}

func tt(x interface{}, y interface{}) bool {

	v := reflect.TypeOf(x)
	w := reflect.TypeOf(y)
	return v == w
}
