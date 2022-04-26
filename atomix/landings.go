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

// CleanLandingsAll cuts out redundant moves in all solutions.
func CleanLandingsAll() {
	count := 0
	for _, s := range Solutions[:] {
		landings := AtomLandings(s)
		saved := CleanLandings(landings)
		if saved == 0 {
			count++
			continue
		}
		fmt.Printf("%d=== %s %s:%s===\n", saved, s.UID, s.LevelSet, s.ID)
	}
	fmt.Println("no change", count)

}

// CleanLandings cuts out redundant moves in a solution.
func CleanLandings(lands []AtomList) int {
	end := len(lands)
	inputLength := end
	i := -1
	for {
		i++
		want := lands[i]
		j := end
		for {
			j--
			if want.Equal(lands[j]) {
				copy(lands[i:], lands[j:end])
				end = end - (j - i)
				break
			}
			if j <= i+1 {
				break
			}
		}
		if i > end-3 {
			break
		}
	}
	return inputLength - end
}

// AtomListPrint prints an atomList one atom per line.
func alp(vv []AtomList) {
	for i, v := range vv {
		fmt.Println(i, v)
	}
}

// FindMatchNext finds repeated landings.
func FindMatchNext(lands []AtomList) int {

	for j := len(lands) - 1; j > 0; j-- {
		//if reflect.DeepEqual(lands[0], lands[j]) {
		if lands[0].Equal(lands[j]) {
			return j
		}
	}
	return 0
}
