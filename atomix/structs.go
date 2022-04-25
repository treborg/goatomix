package atomix

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
