package atomix

import (
	"fmt"
)

// Landings returns a list of landings produced by applying move to grid;
func Landings(s Solution) []Arena {
	h := s.History.HistoryMoves()
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
	h := s.History.HistoryMoves()
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
		_, lands := CleanLandings(landings)
		nCuts := len(landings) - len(lands)
		if nCuts == 0 {
			count++
			continue
		}
		fmt.Printf("%d=== %s %s:%s===\n", nCuts, s.UID, s.LevelSet, s.ID)
	}
	fmt.Println("cuts:", count)

}

/*// CheckCleanLandings checks atom landings produce a valid solution.
func CheckCleanLandings(s){
	grid := s.GetArena()
	h := s.HistoryMoves()
	landings := AtomLandings(s)

	index, lands := CleanLandings(landings)
	nhl := []Moves{}
	for _, p:= range index
		nhl = append(nhl, lands[p])
	}
	sHistory := nhl.ToHistory()
	sHistory.CheckHistory(

}

//*/

// CleanLandings cuts out redundant moves in a solution.
func CleanLandings(landings []AtomList) ([]int, []AtomList) {
	index := make([]int, len(landings))
	lands := make([]AtomList, len(landings))
	for i, atoms := range landings {
		index[i] = i
		lands[i] = atoms
	}
	end := len(lands)
	i := -1
	for {
		i++
		want := lands[i]
		j := end
		for {
			j--
			if want.Equal(lands[j]) {
				copy(index[i:], index[j:end])
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
	return index[:end], lands[:end]
}

// AtomListPrint prints an atomList one atom per line.
func alp(vv []AtomList) {
	for i, v := range vv {
		fmt.Println(i, v)
	}
}
