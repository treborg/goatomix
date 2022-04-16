package main

import (
	"fmt"

	"github.com/treborg/goatomix/levelsets"
	"github.com/treborg/goatomix/sols"
)

func main() {
	for k, v := range levelsets.Sets {
		fmt.Printf("%s: %#v\n", k, len(v.Levels))
	}

	fmt.Println(" ")

	solutions := sols.Solutions
	fmt.Printf("%#v\n", solutions[1])

}
