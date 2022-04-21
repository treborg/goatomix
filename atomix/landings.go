package atomix

import "fmt"

// Landings returns a list of landings produced by applying move to grid;
func Landings(s Solution) {
	h := s.History
	results := make([]Arena, len(h)+1)
	grid := GetArena(s.LevelSet, s.ID)

	fmt.Printf("\n%v\n", grid.String())
	results[0] = grid
	for i, m := range s.History.HistoryList() {
		grid = m.ApplyMove(grid)
		fmt.Printf("\n%v\n", grid.String())

		results[i+1] = grid
	}

}
