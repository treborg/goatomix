package atomix

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomLandings(t *testing.T) {

	assert := assert.New(t)

	err := LoadAllLevels("../levels")
	assert.Nil(err, "cant load files")

	solutions, err := LoadRawSolutions("../sols/solutions.json")
	if err != nil {
		t.Fatal(err)
	}
	assert.True(len(solutions) > 1000)
	s := solutions[0]

	t.Run("AtomLandingsMatchLandings", func(t *testing.T) {

		landings := Landings(s)
		atomList := AtomLandings(s)

		assert.Equal(len(landings), len(s.History)/4+1,
			"landings don't match History")
		assert.Equal(len(atomList), len(landings),
			"atom landings don't match History")

		arena := s.GetArena()
		for i, a := range landings {
			t.Run("Arena should match.", func(t *testing.T) {
				arena.ApplyAtoms(atomList[i])
				if !reflect.DeepEqual(a, arena) {
					t.Error("arena from atomList does not match with landings")
				}
			})
		}
	})
}
