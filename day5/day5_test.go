package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGetDestinationFromSource(t *testing.T) {
	tt := []struct {
		source   int
		mappings []Mapping
		expected int
	}{
		{
			source: 53,
			mappings: []Mapping{
				{destinationRangeStart: 50, sourceRangeStart: 98, rangeLength: 2},
				{destinationRangeStart: 52, sourceRangeStart: 50, rangeLength: 48}},
			expected: 55,
		},
		{
			source: 50,
			mappings: []Mapping{
				{destinationRangeStart: 50, sourceRangeStart: 98, rangeLength: 2},
				{destinationRangeStart: 52, sourceRangeStart: 50, rangeLength: 48}},
			expected: 52,
		},
		{
			source: 98,
			mappings: []Mapping{
				{destinationRangeStart: 50, sourceRangeStart: 98, rangeLength: 2},
				{destinationRangeStart: 52, sourceRangeStart: 50, rangeLength: 48}},
			expected: 50,
		},
		{
			source: 82,
			mappings: []Mapping{
				{destinationRangeStart: 50, sourceRangeStart: 98, rangeLength: 2},
				{destinationRangeStart: 52, sourceRangeStart: 50, rangeLength: 48}},
			expected: 84,
		},
		{source: 84,
			mappings: []Mapping{
				{destinationRangeStart: 0, sourceRangeStart: 15, rangeLength: 37},
				{destinationRangeStart: 37, sourceRangeStart: 52, rangeLength: 2},
				{destinationRangeStart: 39, sourceRangeStart: 0, rangeLength: 15},
			},
			expected: 84,
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d => %+v", tc.source, tc.mappings), func(t *testing.T) {
			result := getDestinationValueFromSource(tc.source, tc.mappings)
			if result != tc.expected {
				t.Errorf("%s: want: %d, got: %d\n", fmt.Sprintf("%d => %+v", tc.source, tc.mappings), tc.expected, result)
			}
		})
	}
}

func TestSolvePartOne(t *testing.T) {
	expected := 35

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
	expected := 46

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
