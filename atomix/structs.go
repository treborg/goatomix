package atomix

import (
	"sort"
)

// LevelSetMap is a map of LevelSet indexed by Name
type LevelSetMap map[string]LevelSet

// LevelSet - a struct to hold a levelset.
type LevelSet struct {
	Name    string   `json:"name"`
	Credit  string   `json:"credit"`
	License string   `json:"license"`
	Levels  []*Level `json:"levels"`
}

// Level - a struct to hold a level.
type Level struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Formula string `json:"formula"`

	ArenaS    []string `json:"arena"`
	MoleculeS []string `json:"molecule"`

	Atoms map[string][]string `json:"atoms"`

	Order    int
	Arena    Arena
	Molecule Molecule
	AtomList AtomList
}

// Molecule represents a Levels molecule
type Molecule [][]byte

// AtomPos describes the type and position of an atom in an arena.
type AtomPos struct {
	A byte
	R byte
	C byte
}

// AtomList is a list of all atoms and ther positions in an arena
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
