package levelsets

import (
	"encoding/json"
	"fmt"
	"os"
)

// Read a JSON file
func Read(fn string) (*LevelSet, error) {

	defer func() {
		if err := recover(); err != nil {
			msg := "Unexpected Error: levelsets.Read(" + fn + "):\n "
			fmt.Println(msg, err)
		}
	}()

	var (
		err      error     = nil
		ord      int       = -1
		levelSet *LevelSet = &LevelSet{}
		errFmt   string    = "'%s', Level: %d, Key: '%s', wanted string got %T"
	)

	// start utility functions
	toString := func(k string, v interface{}) (string, error) {
		if !tt(v, "") {
			return "", fmt.Errorf(errFmt, fn, ord, k, v)
		}
		return v.(string), nil
	}

	toGrid := func(k string, v interface{}) ([][]byte, error) {

		grid := [][]byte{}

		if tt(v, []interface{}{}) {
			return grid, fmt.Errorf(
				"'%s', Level: %d, Key: '%s', wanted string got %T",
				fn, ord, k, v,
			)
		}
		v = []string(v.([]string))
		fmt.Printf("%s, type: %T\n", v, v)
		/*
			for _, line := range v {
				b := []byte(line.(string))
				grid = append(grid, b)
			}
			//*/
		return grid, nil
	}
	// end utility

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
			levelSet.Name, err = toString(k, v)
		case "credits":
			levelSet.Credit, err = toString(k, v)
		case "license":
			levelSet.License, err = toString(k, v)
		}
		if err != nil {
			return levelSet, err
		}
	}

	levels := jm["levels"].([]interface{})
	oLevels := []Level{}

	ord = -1
	for _, level := range levels {
		ord++
		oLevel := &Level{}

		for k, v := range level.(map[string](interface{})) {
			switch k {

			case "name":
				oLevel.Name, err = toString(k, v)
			case "id":
				oLevel.ID, err = toString(k, v)

			case "formula":
				oLevel.Formula, err = toString(k, v)

			case "arena":
				oLevel.Arena, err = toGrid(k, v)
			case "molecule":
				oLevel.Molecule, err = toGrid(k, v)
			}
			if err != nil {
				return levelSet, err
			}
		}
		oLevels = append(oLevels, *oLevel)
	}
	levelSet.Levels = oLevels

	return levelSet, nil
}
