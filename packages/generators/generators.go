package generators

import (
	"fmt"
	"os"

	"github.com/epicmeme/aoc/packages/input"
)

func Generate(year int, day int) error {
	fmt.Printf("Generating for %d/%d\n", year, day)

	// Create the directory
	os.MkdirAll(fmt.Sprintf("%d/day%02d", year, day), 0755)

	// Fetch the input
	input, err := input.HttpFetchInput(year, day)
	if err != nil {
		return err
	}
	os.WriteFile(fmt.Sprintf("%d/day%02d/input.txt", year, day), input, 0644)

	// Create the main.go file
	os.WriteFile(fmt.Sprintf("%d/day%02d/main.go", year, day), []byte(fmt.Sprintf(`
	package main

	import (
		"fmt"

		"github.com/epicmeme/aoc/packages/input"
	)

	func task1(input []byte) {
		fmt.Println("Task 1")
	}

	func task2(input []byte) {
		fmt.Println("Task 2")
	}

	func main() {
		input, err := input.HttpFetchInput(%d, %d)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(input))

		task1(input)
		task2(input)
	}
	`, year, day)), 0644)

	fmt.Printf("Generated %d/%d\n", year, day)

	// Generate benchmark
	os.WriteFile(fmt.Sprintf("%d/day%02d/main_test.go", year, day), []byte(`
	package main

	import (
		"testing"
		"os"
	)

	func BenchmarkTask1(b *testing.B) {
		input, err := os.ReadFile("input.txt")
		if err != nil {
			b.Fatal(err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			task1(input)
		}
	}

	func BenchmarkTask2(b *testing.B) {
		input, err := os.ReadFile("input.txt")
		if err != nil {
			b.Fatal(err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			task2(input)
		}
	}
	`), 0644)

	return nil
}
