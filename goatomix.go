package main

import (
	"fmt"

	"github.com/treborg/goatomix/atomix"
)

func main() {
	for k, v := range atomix.Sets {
		fmt.Printf("%s: %#v\n", k, len(v.Levels))
	}
	fmt.Println(len(atomix.Solutions))

	fmt.Println("=================")
}

func init() {
	atomix.LoadAllLevels()
	atomix.LoadAllSolutions()
}
