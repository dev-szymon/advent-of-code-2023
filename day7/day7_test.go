package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCompareHands(t *testing.T) {
	tt := []struct {
		a          Hand
		b          Hand
		withJokers bool
		expected   int
	}{
		{
			a: Hand{
				cards: "KK677",
				bid:   28,
				figures: map[rune]int{
					'K': 2,
					'6': 1,
					'7': 2,
				},
			},
			b: Hand{
				cards: "KTJJT",
				bid:   220,
				figures: map[rune]int{
					'T': 2,
					'K': 1,
					'J': 2,
				},
			},
			withJokers: false,
			expected:   1,
		},
	}

	for _, tc := range tt {
		name := fmt.Sprintf("%s vs %s", tc.a.cards, tc.b.cards)
		t.Run(name, func(t *testing.T) {
			result := compareHands(tc.a, tc.b, tc.withJokers)
			if result != tc.expected {
				t.Errorf("%s: want: %d, got: %d\n", name, tc.expected, result)
			}
		})
	}
}
func TestSolvePartOne(t *testing.T) {
	expected := 6440

	f, err := os.Open("./input_test.txt")
	if err != nil {
		t.Fatalf("day7: part1: could not open input test file: %+v\n", err)
	}

	result, err := solvePartOne(f)
	if err != nil {
		t.Fatalf("day7: part1: error solivng part one: %+v\n", err)
	}
	if result != expected {
		t.Errorf("day7: part1: want: %d, got: %d\n", expected, result)
	}
}

func TestSolvePartTwo(t *testing.T) {
	expected := 5905

	f, err := os.Open("./input_test.txt")
	if err != nil {
		t.Fatalf("day7: part2: error opening input test file: %+v\n", err)
	}

	result, err := solvePartTwo(f)
	if err != nil {
		t.Fatalf("day7: part2: error solivng part one: %+v\n", err)
	}
	if result != expected {
		t.Errorf("day7: part2: want: %d, got: %d\n", expected, result)
	}
}
