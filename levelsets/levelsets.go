package levelsets

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
	return level.Arena
}

// LoadAll avilable levelsets
func LoadAll() (LevelSetMap, error) {
	names := []string{
		"katomic", "original", "pack1", "mystery", "draknek",
	}

	for _, v := range names {
		fn := fmt.Sprintf("levels/%s.json", v)
		levelSet, err := Load(fn)
		if err != nil {
			return Sets, err
		}
		Sets[levelSet.Name] = levelSet
	}

	return Sets, nil
}

// Load a json file
func Load(fn string) (LevelSet, error) {
	var err error = nil
	file, _ := ioutil.ReadFile(fn)
	set := LevelSet{}

	_ = json.Unmarshal([]byte(file), &set)

	for i, level := range set.Levels {

		level.Order = i + 1

		a, okArena := bytesToSlice(level.ArenaS)

		level.Arena = Arena(a)

		m, okMolecule := bytesToSlice(level.MoleculeS)
		level.Molecule = Molecule(m)

		if !(okArena && okMolecule) {
			err = fmt.Errorf("Loading(%s) level: %d, rows in arenas or molecules must have the same length", fn, level.Order)
		}
		set.Levels[i] = level

		key := set.Name + "!" + level.ID
		LevelMap[key] = set.Levels[i]

	}
	return set, err
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
