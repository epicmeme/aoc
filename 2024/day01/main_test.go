package main

import (
	"os"
	"testing"
)

func BenchmarkTask1(b *testing.B) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		b.Fatal(err)
	}
	col1, col2 := parseInput(input)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		task1(col1, col2)
	}
}

func BenchmarkTask2(b *testing.B) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		b.Fatal(err)
	}
	col1, col2 := parseInput(input)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		task2(col1, col2)
	}

}
