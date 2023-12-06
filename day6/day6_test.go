package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCanBeatCurrentRecord(t *testing.T) {
	tt := []struct {
		time     int
		distance int
		speed    int
		expected bool
	}{
		{
			time:     7,
			distance: 9,
			speed:    1,
			expected: false,
		},
		{
			time:     7,
			distance: 9,
			speed:    2,
			expected: true,
		},
	}

	for _, tc := range tt {
		name := fmt.Sprintf("time: %d, distance %d", tc.time, tc.distance)
		t.Run(name, func(t *testing.T) {
			result := canBeatCurrentRecord(tc.speed, tc.time, tc.distance)
			if result != tc.expected {
				t.Errorf("%s: want: %v, got: %v", name, tc.expected, result)
			}
		})
	}
}
func TestSolvePartOne(t *testing.T) {
	expected := 288

	f, err := os.Open("./input_test.txt")
	if err != nil {
		t.Fatalf("day5: part1: could not open input test file: %+v\n", err)
	}

	result, err := solvePartOne(f)
	if err != nil {
		t.Fatalf("day5: part1: error solivng part one: %+v\n", err)
	}
	if result != expected {
		t.Errorf("day5: part1: want: %d, got: %d\n", expected, result)
	}
}

func TestSolvePartTwo(t *testing.T) {
	expected := 71503

	f, err := os.Open("./input_test.txt")
	if err != nil {
		t.Fatalf("day5: part2: error opening input test file: %+v\n", err)
	}

	result, err := solvePartTwo(f)
	if err != nil {
		t.Fatalf("day5: part2: error solivng part one: %+v\n", err)
	}
	if result != expected {
		t.Errorf("day5: part2: want: %d, got: %d\n", expected, result)
	}
}
