package main

import (
	"fmt"
	"log"

	"github.com/treborg/goatomix/atomix"
)

func main() {
	for k, v := range atomix.Sets {
		fmt.Printf("%s: %#v\n", k, len(v.Levels))

	}
	fmt.Println(len(atomix.Solutions))

	fmt.Println("=================")

	arena := atomix.GetArena("katomic", "83")
	atoms := atomix.ScanGrid(arena)
	fmt.Println("atoms,", atoms)

	fmt.Println(arena.String())
	arena.Clear()
	fmt.Println(arena.String())
}

func init() {
	_, err := atomix.LoadAllLevels()
	if err != nil {
		log.Fatal(err)
	}
	err = atomix.LoadAllSolutions()
	if err != nil {
		log.Fatal(err)
	}
}
