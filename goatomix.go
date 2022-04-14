package main

import (
	"fmt"
	"log"

	"github.com/treborg/goatomix/levelsets"
)

func main() {
	ls, err := levelsets.Read("levels/draknek.json")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println(ls.Name)

	fmt.Println(len(ls.Levels))
	l := ls.Levels[1]
	fmt.Printf("level:1 %T\n", l)

}
