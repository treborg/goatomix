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
	solutions := atomix.GetSolutions()
	fmt.Println(len(solutions))

	err := atomix.WriteSolutions("sols/solutions.jsonl", solutions)
	if err != nil {
		log.Fatal(err)
	}
	//s := atomix.Solutions[2000]

	//tems := atomix.Landings(atomix.Solutions[0])
	//	_ = atomix.AtomLandings()
	//s = atomix.FindSolution("fHhLnZxP")
	fmt.Println("Cleaning histories")

	atomix.CleanHistoryAll()
}

func init() {
	err := atomix.LoadAllLevels("levels")
	if err != nil {
		log.Fatal(err)
	}
	err = atomix.LoadSolutions("sols/solutions.json")
	if err != nil {
		log.Fatal(err)
	}
}
