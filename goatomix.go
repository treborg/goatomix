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

	atoms := atomix.ScanGrid(atomix.GetArena("katomic", "23"))
	fmt.Println("atoms,", atoms)
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
