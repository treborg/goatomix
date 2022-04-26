package atomix

import (
	"fmt"
)

// Move holds the start and end position of a move.
type Move struct {
	SC byte // start column
	SR byte // start row
	EC byte // end row
	ER byte // end column
}

// NewMove creates a new move struct from a string.
func NewMove(sMove string) Move {
	m := Move{}
	m.FromString(sMove)
	return m
}

// FromString set the Move from a four letter string.
func (m *Move) FromString(s string) {
	if len(s) != 4 {
		return
	}
	m.SR = byte(s[0]) - 'a'
	m.SC = byte(s[1]) - 'a'
	m.ER = byte(s[2]) - 'a'
	m.EC = byte(s[3]) - 'a'
}

const (
	// EMPTY represents an empty square in an an Arena
	EMPTY = '.'
	// WALL represents a square containig a wall an Arena.
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

//CheckMove - Check if a move is valid when applied to an arena.
func (m Move) CheckMove(grid Arena) error {

	var w, h byte
	h = byte(len(grid) - 1)
	if (h + 1) > 3 {
		w = byte(len(grid[0]) - 1)
	}
	if (w + 1) < 4 {
		return fmt.Errorf("checkMove grid too small")
	}

	// Make sure an erronous move doesnt cause a panic.
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
