package levelsets

import (
	"encoding/json"
	"os"
)

// Read a JSON file
func Read(fn string) (*LevelSet, error) {

	levelSet := &LevelSet{}

	bytes, err := os.ReadFile(fn)
	if err != nil {
		return levelSet, err
	}

	jm := make(map[string](interface{}))
	err = json.Unmarshal(bytes, &jm)
	if err != nil {
		return levelSet, err
	}

	for k, v := range jm {
		switch k {
		case "name":
			levelSet.Name = v.(string)
		case "credits":
			levelSet.Credit = v.(string)
		case "license":
			levelSet.License = v.(string)
		}
	}

	levels := jm["levels"].([]interface{})
	oLevels := []Level{}
	for _, level := range levels {

		oLevel := &Level{}

		for k, v := range level.(map[string](interface{})) {

			switch k {

			case "name":
				oLevel.Name = v.(string)
			case "id":
				oLevel.ID = v.(string)
			case "formula":
				oLevel.Formula = v.(string)
			case "arena":
				oLevel.Arena = grid(v.([]interface{}))
			case "molecule":
				oLevel.Molecule = grid(v.([]interface{}))
			default:
			}

		}
		oLevels = append(oLevels, *oLevel)
	}
	levelSet.Levels = oLevels

	return levelSet, nil
}

func grid(v []interface{}) [][]byte {
	grid := [][]byte{}
	for _, line := range v {
		b := []byte(line.(string))
		grid = append(grid, b)
	}
	return grid
}

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
