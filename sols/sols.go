package sols

import (
	"encoding/json"
	"os"
)

// Solutions is a list of solutions.
type Solutions []Solution

// Solution holds a solution.
type Solution struct {
	UID      string `json:"uid"`
	Date     string `json:"date"`
	LevelSet string `json:"levelSet"`
	ID       string `json:"id"`
	User     string `json:"user"`
	History  string `json:"history"`
}

// Read a json string
func Read(fn string) (Solutions, error) {
	var err error = nil
	solutions := Solutions{}
	file, err := os.ReadFile(fn)
	if err != nil {
		return solutions, err
	}

	err = json.Unmarshal([]byte(file), &solutions)
	if err != nil {
		return solutions, err
	}

	return solutions, err
}
