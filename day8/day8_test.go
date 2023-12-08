package main

import (
	"os"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	expected := 6

	f, err := os.Open("./input_test.txt")
	if err != nil {
		t.Fatalf("day8: part1: could not open input test file: %+v\n", err)
	}

	result := solvePartOne(f)
	if result != expected {
		t.Errorf("day8: part1: want: %d, got: %d\n", expected, result)
	}
}

func TestSolvePartTwo(t *testing.T) {
	expected := 6

	f, err := os.Open("./input_test2.txt")
	if err != nil {
		t.Fatalf("day8: part2: error opening input test file: %+v\n", err)
	}

	result := solvePartTwo(f)
	if result != expected {
		t.Errorf("day8: part2: want: %d, got: %d\n", expected, result)
	}
}
