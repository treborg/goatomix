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

	bh := []byte(h)
	for i, v := range bh {
		bh[i] = v - 'a'
	}
	for i := 0; i < len(bh); i += 4 {
		fmt.Printf("%s\n", h[i:i+4])

		fmt.Printf("ch: i=%d\n", i)

		sc, sr, ec, er := bh[i], bh[i+1], bh[i+2], bh[i+3]
		fmt.Printf("Move sr.%d er.%d sc.%d ec.%d\n", sr, er, sc, ec)

		err := CheckMove(grid, sc, sr, ec, er)
		if err != nil {
			return err
		}
		grid[er][ec] = grid[sr][sc]
		grid[sr][sc] = EMPTY
		fmt.Println("ch move done")

	}

	return nil
}
