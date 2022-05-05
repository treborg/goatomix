package atomix

import (
	"bytes"
	"sort"
)

// Arena represent a Levels arena
type Arena [][]byte

// Show print arena
func (a Arena) String() string {
	return string(bytes.Join(a, []byte("\n")))
}

// Copy method returns a deep copy of the arena.
func (a Arena) Copy() Arena {
	newRows := make(Arena, len(a))
	for i, row := range a {
		newRow := make([]byte, len(row))
		copy(newRow, row)
		newRows[i] = newRow
	}
	return newRows
}

// Clear removes all atoms from an arena (in place).
func (a Arena) Clear() {
	for r, row := range a {
		for c, sq := range row {
			if isAtom(sq) {
				(a)[r][c] = EMPTY
			}
		}
	}
}

// ApplyMove applies a move to an Arena in place.
// No checks are made, invalid moves may cause a panic.
func (a Arena) ApplyMove(m Move) {
	a[m.ER][m.EC], a[m.SR][m.SC] = a[m.SR][m.SC], EMPTY
}

// ApplyAtoms from an AtomList to a clean areana (in place)
// No checks are made, invalid moves may cause a panic.
func (a Arena) ApplyAtoms(atoms AtomList) {
	a.Clear()
	for _, atom := range atoms {
		a[atom.R][atom.C] = atom.A
	}
}

// FindAtoms returns a sorted list of atoms and their position.
func (a Arena) FindAtoms() AtomList {

	atoms := AtomList{}
	for r, row := range a {
		for c, sq := range row {
			if !isAtom(sq) {
				continue
			}
			atom := AtomPos{sq, byte(r), byte(c)}
			atoms = append(atoms, atom)
		}
	}
	sort.Sort(atoms)

	return atoms
}
