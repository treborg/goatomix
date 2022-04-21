package atomix

import (
	"fmt"
)

// Move is a structure that holds a move.
type Move struct {
	Move string
	SC   byte
	SR   byte
	EC   byte
	ER   byte
}

// ApplyMove applies a move to a grid.
func (m Move) ApplyMove(grid Arena) Arena {
	grid[m.ER][m.EC], grid[m.SR][m.SC] = grid[m.SR][m.SC], EMPTY
	return grid
}

// NewMove creates a new move struct from a string.
func NewMove(move string) Move {
	m := Move{
		Move: move,
		SR:   byte(move[0]) - 'a', // start row
		SC:   byte(move[1]) - 'a', // start column
		ER:   byte(move[2]) - 'a', // end row
		EC:   byte(move[3]) - 'a', // end column
	}
	return m
}

const (
	// EMPTY represents an empty square in the grid.
	EMPTY = '.'
	// WALL represents a square containig a wall in the grid.
	WALL = '#'
)

func isAtom(c byte) bool {
	return !isWall(c) && !isEmpty(c)
}
func isWall(c byte) bool {
	return c == WALL
}
func isEmpty(c byte) bool {
	return c == EMPTY
}

//CheckMove - Check if a move is valid when applied to 'grid'.
func (m Move) CheckMove(grid Arena) error {

	// fmt.Printf("Check Move: %+v \n", grid)
	// return fmt.Errorf("check move break")

	w, h := byte(len(grid[0])-1), byte(len(grid)-1)
	if m.SC >= w || m.EC >= w || m.SR >= h || m.ER >= h {
		return fmt.Errorf("bad move, out of bounds")
	}

	if !isAtom(grid[m.SR][m.SC]) {
		return fmt.Errorf("bad move, no atom in starting position")
	}

	if m.SR != m.ER && m.SC != m.EC {
		return fmt.Errorf("bad move, diagonal move")
	}

	if m.SC < m.EC {
		// move right
		c := m.SC
		for c < m.EC {
			c++
			sq := grid[m.SR][c]
			if !isEmpty(sq) {
				return fmt.Errorf("bad move right, move blocked")
			}
		}
		sq := grid[m.SR][c+1]
		if isEmpty(sq) {
			return fmt.Errorf("bad move right, move NOT blocked")
		}
		return nil
	}

	if m.SC > m.EC {
		// move left
		c := m.SC
		for c > m.EC {
			c--
			sq := grid[m.SR][c]
			if !isEmpty(sq) {
				return fmt.Errorf("bad move left, move blocked")
			}
		}
		sq := grid[m.SR][c-1]
		if isEmpty(sq) {
			return fmt.Errorf("bad move left, move NOT blocked")
		}
		return nil
	}

	if m.SR < m.ER {
		// move down
		r := m.SR
		for r < m.ER {
			r++
			sq := grid[r][m.EC]
			if !isEmpty(sq) {
				return fmt.Errorf("bad move down, move blocked")
			}
		}
		sq := grid[r+1][m.EC]
		if isEmpty(sq) {
			return fmt.Errorf("bad move down, move NOT blocked")
		}
		return nil
	}

	if m.SR > m.ER {
		// move up
		r := m.SR
		for r > m.ER {
			r--
			sq := grid[r][m.EC]
			if !isEmpty(sq) {
				return fmt.Errorf("bad move up, move blocked")
			}
		}
		sq := grid[r-1][m.EC]
		if isEmpty(sq) {
			return fmt.Errorf("bad move up, move NOT blocked")
		}
		return nil
	}
	return fmt.Errorf("bad move, no move")
}
