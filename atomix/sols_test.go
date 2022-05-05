package atomix

import "testing"

func TestReadSolutionFromFile(t *testing.T) {
	ss, err := ReadSolutionsFromFile("../sols/solutions.jsonl")

	t.Run("Should run without error", func(t *testing.T) {
		if err != nil {
			t.Error(err)
		}
	})

	//for _, s := range ss {
	//	t.Logf("%+v\n", s)
	//}

	t.Run("should return at least 1 solution", func(t *testing.T) {
		if len(ss) < 1 {
			t.Error("No results")
		}
	})
}
