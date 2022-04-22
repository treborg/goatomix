package atomix

import "bytes"

// Arena represent a Levels arena
type Arena [][]byte

// Show print arena
func (a *Arena) String() string {
	return string(bytes.Join(*a, []byte("\n")))
}

// Copy method returns a deep copy of the arena.
func (a *Arena) Copy() Arena {
	newRows := make([][]byte, len(*a))
	for i, row := range *a {
		newRow := make([]byte, len(row))
		copy(newRow, (*a)[i])
		newRows[i] = newRow
	}
	return Arena(newRows)
}

// Clear removes all atoms from an arena.
func (a *Arena) Clear() {
	for r, row := range *a {
		for c, sq := range row {
			if isAtom(sq) {
				(*a)[r][c] = EMPTY
			}
		}
	}
}

// ApplyAtoms cleans the arena and populates it with atoms from an AtomList.
func (a *Arena) ApplyAtoms(atoms AtomList) {
	a.Clear()
	for _, atom := range atoms {
		(*a)[atom.R][atom.C] = atom.A
	}
}
