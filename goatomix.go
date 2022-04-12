package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	bytes, err := os.ReadFile("levels/draknek.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))

}
