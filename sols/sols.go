package sols

import (
	"encoding/json"
	"os"
)

// Solutions is a variable holding a list of Solution structs.
var Solutions = SolutionList{}

func init() {
	sols, err := Load("sols/solutions.json")
	if err != nil {
		panic(err)
	}
	Solutions = sols
}

// SolutionList is a list of solutions.
type SolutionList []Solution

// Solution holds a solution.
type Solution struct {
	UID      string `json:"uid"`
	Date     string `json:"date"`
	LevelSet string `json:"levelSet"`
	ID       string `json:"id"`
	User     string `json:"user"`
	History  string `json:"history"`
}

// Load a json string
func Load(fn string) (SolutionList, error) {
	var err error = nil
	solutions := SolutionList{}
	sols, err := os.ReadFile(fn)
	if err != nil {
		return solutions, err
	}

	err = json.Unmarshal([]byte(sols), &solutions)
	if err != nil {
		return solutions, err
	}

	return solutions, err
}
