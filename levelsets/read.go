package levelsets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Read a json string
func Read(fn string) {
	file, _ := ioutil.ReadFile(fn)
	data := LevelSet{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.Levels); i++ {
		fmt.Println("Name: ", data.Levels[i].Name)
		fmt.Println("Id: ", data.Levels[i].ID)
	}
}
