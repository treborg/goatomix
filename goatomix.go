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
	//s := atomix.Solutions[2000]

	//tems := atomix.Landings(atomix.Solutions[0])
	//	_ = atomix.AtomLandings()
	//s = atomix.FindSolution("fHhLnZxP")
	atomix.CleanLandingsAll()
}

func init() {
	err := atomix.LoadAllLevels("levels")
	if err != nil {
		log.Fatal(err)
	}
	err = atomix.LoadAllSolutions("sols/solutions.json")
	if err != nil {
		log.Fatal(err)
	}
}
