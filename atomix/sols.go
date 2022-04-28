package atomix

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
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

// NewSolution returns a new Solution with minimal setup.
func NewSolution(levelSet, id string, history History) Solution {
	s := Solution{
		LevelSet: levelSet,
		ID:       id,
		History:  history,
	}
	return s
}

// GetSolutions returns the list of solutions.
func GetSolutions() SolutionList {
	return Solutions
}

// ToJSON converts a solution a json object.
func (s *Solution) ToJSON() {
	// fmt.Printf("solution %+v\n", s)
	j, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("json", string(j))
	fmt.Println("json", string(j))

}

// CheckHistory for this solutions is valid.
func (s *Solution) CheckHistory() error {
	grid := s.GetArena()
	err := s.History.CheckHistory(grid)
	if err != nil {
		return fmt.Errorf("error in solution %s: %s ", s.UID, err)
	}
	return nil
}

// ToMoveList returns this solutions History as a []Moves.
func (s *Solution) ToMoveList() MoveList {
	return s.History.ToMoveList()
}

// GetArena the starting Arena for this solutions level.
func (s *Solution) GetArena() Arena {
	return GetArena(s.LevelSet, s.ID)
}

// LoadSolutions from a file and check the validity of
// each solutions History.
func LoadSolutions(fn string) error {

	solutions, err := LoadRawSolutions(fn)
	if err != nil {
		return err
	}
	Solutions = solutions
	for i, s := range solutions {
		err := s.CheckHistory()

		if err != nil {
			err := fmt.Errorf("invalid solution %d from '%s', %s", i, fn, err)
			return err
		}
	}
	return nil
}

// LoadRawSolutions from a json file.
func LoadRawSolutions(fn string) (SolutionList, error) {
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

	fmt.Println("Load", len(solutions))

	return solutions, err
}

// WriteSolutions writes List of solutions in jsonl style
func WriteSolutions(fn string, ss SolutionList) error {

	file, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, s := range ss {
		bytes, err := json.Marshal(&s)
		if err == nil {
			_, err = writer.Write(bytes)
			if err == nil {
				err = writer.WriteByte('\n')
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

//*/
