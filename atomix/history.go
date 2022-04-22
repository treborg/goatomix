package atomix

import (
	"fmt"
)

// History container for history.
type History string

// HistoryList method converts History to []Move.
func (h History) HistoryList() []Move {
	sh := string(h)
	moves := make([]Move, 0, len(h)/4)
	for i := 0; i < len(h); i += 4 {
		move := NewMove(sh[i : i+4])
		moves = append(moves, move)
	}
	return moves
}

// CheckHistory is a method to verify that this history is valid.
func (h History) CheckHistory(grid Arena) error {

	grid = grid.Copy()

	if len(h)%4 != 0 {
		msg := "history length %d:  must be multiple of 4 chars: %s"
		return fmt.Errorf(msg, len(h), h)
	}

	for _, m := range h.HistoryList() {
		err := m.CheckMove(grid)
		if err != nil {
			return err
		}
		grid.ApplyMove(m)
	}
	return nil
}
