package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func task1(col1 []int, col2 []int) int {
	var sortedCol1 = slices.Clone(col1)
	slices.Sort(sortedCol1)

	var sortedCol2 = slices.Clone(col2)
	slices.Sort(sortedCol2)

	var tot = 0

	for i := 0; i < len(sortedCol1); i++ {
		var delta = sortedCol1[i] - sortedCol2[i]
		if delta < 0 {
			tot += -delta
		} else {
			tot += delta
		}
	}

	return tot
}

func task2(col1 []int, col2 []int) int {
	var sortedCol1 = slices.Clone(col1)
	slices.Sort(sortedCol1)

	var sortedCol2 = slices.Clone(col2)
	slices.Sort(sortedCol2)

	var totalSimilarity = 0

	for i := 0; i < len(sortedCol1); i++ {
		var freq = 0
		for j := 0; j < len(sortedCol2); j++ {
			if sortedCol1[i] == sortedCol2[j] {
				freq++
			}
		}
		totalSimilarity += freq * sortedCol1[i]
	}

	return totalSimilarity
}

func parseInput(input []byte) ([]int, []int) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	var col1 []int
	var col2 []int

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")

		num1, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		col1 = append(col1, num1)

		num2, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		col2 = append(col2, num2)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return col1, col2
}

func main() {
	body, err := os.ReadFile("2024/day01/input.txt")
	if err != nil {
		panic(err)
	}

	col1, col2 := parseInput(body)

	fmt.Println("Running tasks...")
	fmt.Printf("Task 1: %v\n", task1(col1, col2))
	fmt.Printf("Task 2: %v\n", task2(col1, col2))
}
