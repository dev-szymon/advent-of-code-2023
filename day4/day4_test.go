package main

import (
	"os"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	expected := 13

	f, err := os.Open("./input_test.txt")
	if err != nil {
		t.Fatalf("day4: error opening input test file: %+v\n", err)
	}
	defer f.Close()

	result, err := solvePartOne(f)
	if err != nil {
		t.Fatalf("day4: error solving part1: %+v\n", err)
	}

	if result != expected {
		t.Errorf("day4, part2: got: %d, want: %d\n", result, expected)
	}
}

func TestSolvePartTwo(t *testing.T) {
	expected := 30

	f, err := os.Open("./input_test.txt")
	if err != nil {
		t.Fatalf("day4: error opening input test file: %+v\n", err)
	}
	defer f.Close()

	result, err := solvePartTwo(f)
	if err != nil {
		t.Fatalf("day4: error solving part2: %+v\n", err)
	}

	if result != expected {
		t.Errorf("day2, part2: got: %d, want: %d\n", result, expected)
	}
}
