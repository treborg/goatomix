package levelsets

import "bytes"

// Arena represent a Levels arena
type Arena [][]byte

// Show print arena
func (a *Arena) String() string {
	return string(bytes.Join(*a, []byte("\n")))
}
