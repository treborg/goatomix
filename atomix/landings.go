package atomix

import (
	"fmt"
	"log"
)

// Landings returns a list of landings as a []Array;
func Landings(s Solution) []Arena {
	h := s.ToMoveList()
	results := make([]Arena, len(h)+1)
	grid := s.GetArena()

	results[0] = grid
	for i, move := range h {
		grid = grid.Copy()
		grid.ApplyMove(move)
		results[i+1] = grid
	}
	return results
}

// AtomLandings returns a list of landings as a []AtomList
func AtomLandings(s Solution) []AtomList {
	h := s.ToMoveList()
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

// CleanHistoryAll cuts out redundant moves for all solutions.
func CleanHistoryAll() {
	count := 0
	sols := GetSolutions()

	fmt.Println("Got solutions ", len(sols))

	for _, s := range sols {
		sLen := len(s.History) / 4
		sClean, err := CleanHistory(s)
		if err != nil {
			log.Fatal(err)
		}
		lenClean := len(sClean.History) / 4
		nCuts := sLen - lenClean
		if nCuts == 0 {
			continue
		}
		count++
		fmt.Printf("%03d %03d === %s %s:%s===\n", nCuts, sLen, s.UID, s.LevelSet, s.ID)
	}
	fmt.Println("cuts:", count)

}

//*/

// CleanHistory removes redundant moves.
func CleanHistory(s Solution) (Solution, error) {

	landings := AtomLandings(s)
	index := CleanLandings(landings)

	moves := s.ToMoveList()
	newMoves := make(MoveList, len(index))

	for i, m := range index {
		newMoves[i] = moves[m]
	}
	s.History = newMoves.ToHistory()

	err := s.CheckHistory()
	if err != nil {
		return s, err
	}
	return s, nil
}

//*/

// CleanLandings cuts out redundant moves in a solution.
func CleanLandings(landings []AtomList) []int {
	index := make([]int, len(landings))
	lands := make([]AtomList, len(landings))
	for i, atoms := range landings {
		index[i] = i
		lands[i] = atoms
	}
	end := len(lands) - 1
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
	return index[:end]
}
