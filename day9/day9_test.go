package main

import (
	"fmt"
	"os"
	"testing"
)

func TestExtrapolate(t *testing.T) {
	tt := []struct {
		input     []int
		expected  int
		direction int
	}{
		{
			input:     []int{1, 3, 6, 10, 15, 21},
			expected:  28,
			direction: forward,
		},
		{
			input:     []int{10, 13, 16, 21, 30, 45},
			expected:  68,
			direction: forward,
		},
		{
			input:     []int{-5, -3, -1, 1, 3, 5},
			expected:  7,
			direction: forward,
		},
		{
			input:     []int{10, 13, 16, 21, 30, 45},
			expected:  5,
			direction: backward,
		},
	}

	for _, tc := range tt {
		name := fmt.Sprintf("%+v", tc.input)
		t.Run(name, func(t *testing.T) {
			result := extrapolateValue(tc.input, tc.direction)
			if result != tc.expected {
				t.Errorf("%s: want: %d, got: %d\n", name, tc.expected, result)
			}
		})
	}
}

func TestSolvePartOne(t *testing.T) {
	expected := 114

	f, err := os.Open("./input_test.txt")
	if err != nil {
		t.Fatalf("day9: part1: could not open input test file: %+v\n", err)
	}

	result, err := solvePartOne(f)
	if err != nil {
		t.Fatalf("day9: part1: error solivng part one: %+v\n", err)
	}
	if result != expected {
		t.Errorf("day9: part1: want: %d, got: %d\n", expected, result)
	}
}

func TestSolvePartTwo(t *testing.T) {
	expected := 2

	f, err := os.Open("./input_test.txt")
	if err != nil {
		t.Fatalf("day9: part1: could not open input test file: %+v\n", err)
	}

	result, err := solvePartTwo(f)
	if err != nil {
		t.Fatalf("day9: part1: error solivng part one: %+v\n", err)
	}
	if result != expected {
		t.Errorf("day9: part1: want: %d, got: %d\n", expected, result)
	}
}
