package atomix

import (
	"encoding/json"
	"fmt"
	"os"
)

// Solutions is a variable holding a list of Solution structs.
var Solutions = SolutionList{}

// SolutionList is a list of solutions.
type SolutionList []Solution

// Solution holds a solution.
type Solution struct {
	UID      string  `json:"uid"`
	Date     string  `json:"date"`
	LevelSet string  `json:"levelSet"`
	ID       string  `json:"id"`
	User     string  `json:"user"`
	History  History `json:"history"`
}

// CheckHistory for this solutions is valid.
func (s Solution) CheckHistory() error {
	grid := GetArena(s.LevelSet, s.ID)
	err := s.History.CheckHistory(grid)
	if err != nil {
		return fmt.Errorf("error in solution %s: %s ", s.UID, err)
	}
	return nil
}

// LoadAllSolutions from solutions.json
func LoadAllSolutions() {

	sols, err := LoadSolutions("sols/solutions.json")
	if err != nil {
		panic(err)
	}
	Solutions = sols
	for _, s := range sols {
		err := s.CheckHistory()

		if err != nil {
			panic(err)
		}
	}
}

// LoadSolutions from a json file.
func LoadSolutions(fn string) (SolutionList, error) {
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
