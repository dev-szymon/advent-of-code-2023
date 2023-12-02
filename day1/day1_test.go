package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	expected := 142
	f, err := os.Open("./input_test1.txt")
	if err != nil {
		t.Fatal("could not open test input file")
	}
	defer f.Close()

	result, err := solvePartOne(f)
	if err != nil {
		t.Fatalf("Error solving day1 part1: %+v", err)
	}

	if result != expected {
		t.Fatalf("got: %d, want: %d", result, expected)
	}
}

func TestFindCalibrationValues(t *testing.T) {
	tt := []struct {
		line      string
		withWords bool
		expected  int
	}{
		{
			line:      "ab3cone2threex83yz",
			withWords: false,
			expected:  33,
		},
		{
			line:      "three2",
			withWords: false,
			expected:  22,
		},
		{
			line:      "abcone2threexyz",
			withWords: true,
			expected:  13,
		},
		{
			line:      "1ninehgqtjprgnpkchxdkctzk",
			withWords: true,
			expected:  19,
		},
		{
			line:      "oneight",
			withWords: true,
			expected:  18,
		},
		{
			line:      "asdf3eight2zeroasdf",
			withWords: true,
			expected:  32,
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%s => %d", tc.line, tc.expected), func(t *testing.T) {
			result, err := findCalibrationValue(tc.line, tc.withWords)
			if err != nil {
				t.Fatalf("error trying to find calibration valur: %+v\n", err)
			}
			if result != tc.expected {
				t.Errorf("got: %d, want: %d", result, tc.expected)
			}
		})
	}

}

func TestSolvePartTwo(t *testing.T) {
	expected := 281
	f, err := os.Open("./input_test2.txt")
	if err != nil {
		t.Fatal("could not open test input file")
	}
	defer f.Close()

	p2, err := solvePartTwo(f)
	if err != nil {
		t.Fatalf("Error solving day1 part2: %+v", err)
	}

	if p2 != expected {
		t.Fatalf("expected: %d, got: %d", expected, p2)
	}
}
