package main

import (
	"os"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	expected := 4361
	f, err := os.Open("./input_test.txt")
	if err != nil {
		t.Fatalf("day3: error opening input test file: %+v\n", err)
	}
	defer f.Close()

	result, err := solvePartOne(f)
	if err != nil {
		t.Errorf("day3: error solving part1: %+v\n", err)
	}

	if result != expected {
		t.Errorf("day3: got: %d, want: %d\n", result, expected)
	}
}

func TestSolvePartTwo(t *testing.T) {
	expected := 467835
	f, err := os.Open("./input_test.txt")
	if err != nil {
		t.Fatalf("day3: error opening input test file: %+v\n", err)
	}
	defer f.Close()

	result, err := solvePartTwo(f)
	if err != nil {
		t.Errorf("day3: error solving part1: %+v\n", err)
	}

	if result != expected {
		t.Errorf("day3: got: %d, want: %d\n", result, expected)
	}
}
