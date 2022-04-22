package atomix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var SArena = []string{
	"#############",
	"#...#..#....#",
	"#.#..#5#....#",
	"#..#...#3.#.#",
	"#......####.#",
	"#.......#1#.#",
	"#.2..#..#...#",
	"#.#####....##",
	"#.4.#..#...#.",
	"#####..#...#.",
	".......#####.",
}

func TestAtomixListCopy(t *testing.T) {
	assert := assert.New(t)
	a, ok := GridToBytes(SArena)
	if !ok {
		t.Fatal()
	}
	arena := Arena(a)

	atoms := ScanGrid(arena)

	xatoms := atoms.Copy()
	xatoms[0] = AtomPos{}
	xatoms[1].R = 99

	assert.Equal(xatoms[0].A, byte(0))
	assert.Equal(xatoms[1].R, byte(99))

	assert.NotEqual(atoms[0].A, byte(0))
	assert.NotEqual(atoms[1].R, byte(99))
}
