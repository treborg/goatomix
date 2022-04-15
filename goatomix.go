package main

import (
	"fmt"

	"github.com/treborg/goatomix/levelsets"
)

func main() {
	levelset, err := levelsets.Read("levels/draknek.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(levelset.Levels))

}
