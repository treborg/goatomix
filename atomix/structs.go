package atomix

// LevelSetMap is a map of LevelSet indexed by Name
type LevelSetMap map[string]LevelSet

// LevelSet a struct to hold a levelset.
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

// AtomPos describes the type and position of an atom in an a grid.
//
//  It is a three byte struct {[type, row, col}
//s
type AtomPos struct {
	A byte
	R byte
	C byte
}

// AtomList is a slice of AtomPos
type AtomList []AtomPos

// Len returns length of slice.
func (a AtomList) Len() int {
	return len(a)
}

// Swap swaps one element with another.
func (a AtomList) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less compares one element with another.
func (a AtomList) Less(i int, j int) bool {

	if a[i].A != a[j].A {
		return a[i].A < a[j].A
	}

	if a[i].R != a[j].R {
		return a[i].R < a[j].R
	}
	return a[i].C < a[i].C
}

// Copy an AtomList.
func (a AtomList) Copy() AtomList {
	newList := make(AtomList, len(a))
	for i, v := range a {
		newList[i] = v
	}
	return newList
}
