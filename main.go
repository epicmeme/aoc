package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/epicmeme/aoc/packages/generators"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: aoc <year> <day>")
		os.Exit(1)
	}
	year, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid year")
		os.Exit(1)
	}
	day, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid day")
		os.Exit(1)
	}
	generators.Generate(year, day)
}
