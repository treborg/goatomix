package atomix

import (
	"fmt"
)

// Landings returns a list of landings produced by applying move to grid;
func Landings(s Solution) []Arena {
	h := s.History.HistoryList()
	results := make([]Arena, len(h)+1)
	grid := GetArena(s.LevelSet, s.ID)

	results[0] = grid
	for i, move := range h {
		grid = grid.Copy()
		grid.ApplyMove(move)
		results[i+1] = grid
	}
	return results
}

// AtomLandings returns a list of landings as an []AtomList
func AtomLandings(s Solution) []AtomList {
	h := s.History.HistoryList()
	results := make([]AtomList, len(h)+1)
	grid := s.GetArena()

	atoms := grid.FindAtoms()

	results[0] = atoms

	for i, move := range h {
		atoms = atoms.Copy()
		atoms.ApplyMove(move)
		results[i+1] = atoms
	}
	return results
}

// FindSolution finds solution by its UID
func FindSolution(uid string) Solution {
	for _, s := range Solutions {
		if s.UID == uid {
			return s
		}
	}
	return Solutions[0]
}

// X does whaever
func X() {
	for _, s := range Solutions[:] {
		FindMatch(s)
	}
}

// FindMatch finds match
func FindMatch(s Solution) {
	fmt.Printf("=== %s %s:%s===\n", s.UID, s.LevelSet, s.ID)

	lands := AtomLandings(s)
	count := 0
	for i := range lands {
		xlands := lands[i:]
		j := FindMatchNext(xlands)
		if j != 0 {
			count++
			//fmt.Println(lands[i])
			//fmt.Println(lands[j+i])

			fmt.Printf("%d, %d\n", i, j+i)
		}
	}

}

// AtomListPrint prints an atomList one atom per line.
func AtomListPrint(vv []AtomList) {
	for i, v := range vv {
		fmt.Println(i, v)
	}
}

// AtomListEqual compares 'other' with
func AtomListEqual(this, other AtomList) bool {
	if len(this) != len(other) {
		return false
	}
	for i, thisPos := range this {
		otherPos := other[i]
		if thisPos.C != otherPos.C ||
			thisPos.R != otherPos.R ||
			thisPos.A != otherPos.A {
			return false
		}
	}
	return true
}

// FindMatchNext finds repeated landings.
func FindMatchNext(lands []AtomList) int {

	for j := len(lands) - 1; j > 0; j-- {
		//if reflect.DeepEqual(lands[0], lands[j]) {
		if AtomListEqual(lands[0], lands[j]) {
			return j
		}
	}
	return 0
}
