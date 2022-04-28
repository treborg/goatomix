package atomix

import (
	"fmt"
)

// History container for history.
type History string

// MoveList is a list of Moves.
type MoveList []Move

// ToHistory turns a list of moves to a History string.
func (hm MoveList) ToHistory() History {
	s := make([]byte, len(hm)*4)
	i := 0
	for _, m := range hm {
		s[i] = m.SR + 'a'
		s[i+1] = m.SC + 'a'
		s[i+2] = m.ER + 'a'
		s[i+3] = m.EC + 'a'
		i += 4
	}
	return History(s)
}

// ToMoveList returns this History to a []Move.
func (h History) ToMoveList() MoveList {
	sh := string(h)
	moves := make([]Move, 0, len(h)/4)
	for i := 0; i < len(h); i += 4 {
		move := NewMove(sh[i : i+4])
		moves = append(moves, move)
	}
	return moves
}

// CheckHistory is a method to verify that this history is valid.
//
// If a sequence of moves which, when applied to the Arena, does not
// produce an error then it is valid, even if it does not end with a solution.
func (h History) CheckHistory(grid Arena) error {

	grid = grid.Copy()

	if len(h)%4 != 0 {
		msg := "history length %d:  must be multiple of 4 chars: %s"
		return fmt.Errorf(msg, len(h), h)
	}

	for _, m := range h.ToMoveList() {
		err := m.CheckMove(grid)
		if err != nil {
			return err
		}
		grid.ApplyMove(m)
	}
	return nil
}
