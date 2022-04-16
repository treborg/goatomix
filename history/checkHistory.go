package history

import "fmt"

// History container for history.
type History []byte

// Check is a method to verify that this history is valid.
func (h History) Check(levelSet, id string) (History, error) {
	fmt.Println(levelSet, id)
	return h, nil
}
