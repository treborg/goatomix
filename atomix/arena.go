package atomix

import (
	"bytes"
	"sort"
)

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

// Clear removes all atoms from an arena (in place).
func (a *Arena) Clear() {
	for r, row := range *a {
		for c, sq := range row {
			if isAtom(sq) {
				(*a)[r][c] = EMPTY
			}
		}
	}
}

// ApplyMove applies a move to a grid in place.
func (a *Arena) ApplyMove(m Move) {
	grid := *a
	grid[m.ER][m.EC], grid[m.SR][m.SC] = grid[m.SR][m.SC], EMPTY
}

// ApplyAtoms from an AtomList to a clean areana (in place)
func (a *Arena) ApplyAtoms(atoms AtomList) {
	a.Clear()
	for _, atom := range atoms {
		(*a)[atom.R][atom.C] = atom.A
	}
}

// ScanGrid finds position of each atom in the grid.
func (a *Arena) ScanGrid() AtomList {

	atoms := AtomList{}
	for r, row := range *a {
		for c, sq := range row {
			if !isAtom(sq) {
				continue
			}
			atom := AtomPos{sq, byte(r), byte(c)}
			atoms = append(atoms, atom)
		}
	}
	sort.Sort(atoms)
	atomList := make(AtomList, len(atoms))
	copy(atomList, atoms)

	return atomList
}
