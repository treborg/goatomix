package levelsets

// LevelSet a struct to hold a levelset.
type LevelSet struct {
	Name    string  `json:"name"`
	Credit  string  `json:"credit"`
	License string  `json:"license"`
	Levels  []Level `json:"levels"`
}

// Level - a struct to hold a level.
type Level struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Formula string `json:"formula"`

	Arena    []string `json:"arena"`
	Molecule []string `json:"molecule"`

	Atoms map[string][]string `json:"atoms"`
}
