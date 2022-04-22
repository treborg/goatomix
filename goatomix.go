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

	xatoms := atoms.Copy()

	xatoms[0] = atomix.AtomPos{}
	xatoms[1].R = 0x99

	fmt.Println("atoms,", atoms[:2])
	fmt.Println("xatoms,", xatoms[:2])

	fmt.Println("=================")

	fmt.Println(arena.String())
	arena.Clear()
	fmt.Println(arena.String())
}

func init() {
	err := atomix.LoadAllLevels()
	if err != nil {
		log.Fatal(err)
	}
	err = atomix.LoadAllSolutions()
	if err != nil {
		log.Fatal(err)
	}
}
