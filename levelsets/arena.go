package levelsets

import "bytes"

// Arena represent a Levels arena
type Arena [][]byte

// Show print arena
func (a Arena) String() string {
	return string(bytes.Join(a, []byte("\n")))
}

// Copy method returns a deep copy of the arena.
func (a Arena) Copy() Arena {
	newRows := make([][]byte, len(a))
	for i, row := range a {
		newRow := make([]byte, len(row))
		for k, c := range row {
			newRow[k] = c
		}
		newRows[i] = newRow
	}
	return Arena(newRows)
}
