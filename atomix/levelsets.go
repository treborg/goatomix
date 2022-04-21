package atomix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Sets is a map of all currently loaded LevelSet
var Sets = LevelSetMap{}

// LevelMap maps levelset name and ids to Level
var LevelMap = map[string]Level{}

// GetLevel returns a level from LevelSet 'name' with ID 'id'
func GetLevel(name, id string) Level {
	return LevelMap[name+"!"+id]
}

// GetArena returns Arena from LevelSet 'name' with ID 'id'
func GetArena(name, id string) Arena {
	key := name + "!" + id
	level, ok := LevelMap[key]
	if !ok {
		panic(fmt.Errorf("No Key, LevelMap[%s!%s]", name, id))
	}
	return level.Arena.Copy()
}

// LoadAllLevels from json levelsets
func LoadAllLevels() (LevelSetMap, error) {
	names := []string{
		"katomic", "original", "pack1", "mystery", "draknek",
	}

	for _, v := range names {
		fn := fmt.Sprintf("levels/%s.json", v)
		levelSet, err := LoadLevels(fn)
		if err != nil {
			return Sets, err
		}
		Sets[levelSet.Name] = levelSet
	}

	return Sets, nil
}

// LoadLevels from a json levelset file.
func LoadLevels(fn string) (LevelSet, error) {
	var err error = nil
	file, _ := ioutil.ReadFile(fn)
	set := LevelSet{}

	_ = json.Unmarshal([]byte(file), &set)

	for i, level := range set.Levels {

		err = fixLevel(i, &level)
		if err != nil {
			return set, fmt.Errorf("file: %s level:%d, %v", fn, i, err)
		}

		key := set.Name + "!" + level.ID
		LevelMap[key] = level
		set.Levels[i] = level

		//fmt.Printf("2 fixed level %d: %+v\n\n",
		//	set.Levels[i].Order,
		//	set.Levels[i].Arena,
		//)

	}
	return set, err
}

// fixLevel
func fixLevel(i int, level *Level) error {

	level.Order = i

	a, okArena := GridToBytes(level.ArenaS)

	level.Arena = Arena(a)

	m, okMolecule := GridToBytes(level.MoleculeS)
	level.Molecule = Molecule(m)

	if !(okArena && okMolecule) {
		err := fmt.Errorf("rows in arenas or molecules must have the same length")
		return err
	}
	return nil
}

// GridToBytes converts grids from []string to [][]byte forms.
func GridToBytes(s []string) ([][]byte, bool) {
	n := len(s[0])
	b := make([][]byte, len(s), len(s))
	ok := true
	for i, r := range s {
		b[i] = []byte(r)
		if n != len(b[i]) {
			ok = false
		}
	}
	return b, ok
}
