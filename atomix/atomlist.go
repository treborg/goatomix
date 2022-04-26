package atomix

import "sort"

// AtomPos describes the type and position of an atom in an arena.
type AtomPos struct {
	A byte // atom type
	R byte // row
	C byte // col
}

// AtomList is a sorted list of all atoms and their positions in an Arena.
type AtomList []AtomPos

// ApplyMove modifies (in place) the positon of the atoms
// in an AtomList and sorts the list after a move.
func (a AtomList) ApplyMove(m Move) {

	for i, atom := range a {
		if atom.R != m.SR || atom.C != m.SC {
			continue
		}
		atom.R = m.ER
		atom.C = m.EC
		a[i] = atom
		break
	}
	sort.Sort(a)
}

// Len returns length of slice.
func (a AtomList) Len() int {
	return len(a)
}

// Swap swaps one element with another.
func (a AtomList) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

// Equal compares two AtomLists for equality.
func (a AtomList) Equal(other AtomList) bool {
	if len(a) != len(other) {
		return false
	}
	for i, thisPos := range a {
		otherPos := other[i]
		if thisPos.C != otherPos.C ||
			thisPos.R != otherPos.R ||
			thisPos.A != otherPos.A {

			return false
		}
	}
	return true
}

// Less compares one element with another.
func (a AtomList) Less(small, big int) bool {

	if a[small].A != a[big].A {
		return a[small].A < a[big].A
	}
	if a[small].R != a[big].R {
		return a[small].R < a[big].R
	}
	return a[small].C < a[big].C
}

// Copy an AtomList.
func (a AtomList) Copy() AtomList {
	newList := make(AtomList, len(a))
	for i, v := range a {
		newList[i] = v
	}
	return newList
}
