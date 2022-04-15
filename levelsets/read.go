package levelsets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Read a json string
func Read(fn string) (LevelSet, error) {
	var err error = nil
	file, _ := ioutil.ReadFile(fn)
	levelset := LevelSet{}

	_ = json.Unmarshal([]byte(file), &levelset)

	for i, level := range levelset.Levels {

		level.Order = i + 1

		a, ok1 := bytesToSlice(level.ArenaS)
		level.Arena = Arena(a)

		m, ok2 := bytesToSlice(level.MoleculeS)
		level.Molecule = Molecule(m)

		if !(ok1 && ok2) {
			err = fmt.Errorf("(%s)level: %d, rows in arenas or molecules must have the same length", fn, level.Order)
		}
	}
	return levelset, err
}

func bytesToSlice(a []string) ([][]byte, bool) {
	n := len(a[0])
	o := make([][]byte, len(a), len(a))
	for i, r := range a {
		if n != len(r) {
			return o, false
		}
		o[i] = []byte(r)
	}
	return o, true
}
