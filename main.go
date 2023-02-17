package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Initiate user input reader
	reader := bufio.NewReader(os.Stdin)

	// Some input tests
	if true {
		var optSystem [3]string
		optSystem[0] = "linux"
		optSystem[1] = "windows"
		optSystem[2] = "mac"

		fmt.Printf("Choose from:\n1. %v\n2. %v\n3. %v\n", optSystem[0], optSystem[1], optSystem[2])
		choiceSystem, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		statement := "%v is %v\n"

		var optRating [3]string
		optRating[0] = "good"
		optRating[1] = "bad"
		optRating[2] = "ugly"

		fmt.Printf("Choose from:\n1. %v\n2. %v\n3. %v\n", optRating[0], optRating[1], optRating[2])
		choiceRating, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		cs, _ := strconv.Atoi(string(choiceSystem[0]))
		cr, _ := strconv.Atoi(string(choiceRating[0]))

		fmt.Printf(statement, optSystem[cs-1], optRating[cr-1])

		return
	}
}
