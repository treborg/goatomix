package history

import (
	"fmt"

	"github.com/treborg/goatomix/levelsets"
)

// History container for history.
type History string

// CheckHistory is a method to verify that this history is valid.
func CheckHistory(name, id string, h History) error {

	grid := levelsets.GetArena(name, id)

	if len(h)%4 != 0 {
		return fmt.Errorf(
			"history length %d:  must be multiple of 4 chars: %s",
			len(h), h,
		)
	}

	sh := string(h)
	for i := 0; i < len(sh); i += 4 {

		m := NewMove(sh[i : i+4])
		err := CheckMove(grid, m)
		if err != nil {
			return err
		}

		grid = m.ApplyMove(grid)
	}
	return nil
}
