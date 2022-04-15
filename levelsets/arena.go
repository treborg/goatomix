package levelsets

import "fmt"

// Arena represent a Levels arena
type Arena [][]byte

// Show print arena
func (a Arena) Show() Arena {
	for _, r := range a {
		fmt.Println(string(r))
	}
	fmt.Println("")

	return a
}
