package main

import (
	"fmt"

	"github.com/treborg/goatomix/atomix"
	"github.com/treborg/goatomix/sols"
)

func main() {
	for k, v := range atomix.Sets {
		fmt.Printf("%s: %#v\n", k, len(v.Levels))
	}
	fmt.Println(len(sols.Solutions))

	fmt.Println("=================")
}

func init() {
	atomix.LoadAllLevels()
	sols.LoadAllSolutions()
}
