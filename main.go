package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	//Ask users for
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg[0])

	if argsWithProg[0] == "--add_car" {

	}

}
