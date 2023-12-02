package main

import (
	"os"
	"testing"
)

func TestCreateGame(t *testing.T) {
	tt := []struct {
		description string
		rounds      []map[string]int
	}{
		{
			description: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			rounds: []map[string]int{
				{"green": 8, "blue": 6, "red": 20},
				{"green": 13, "blue": 5, "red": 4},
				{"green": 5, "red": 1},
			},
		},
		{
			description: "Game 3: 8 green, 6 blue, 2 blue, 20 red; 5 blue, 4 red, 3 blue, 1 red, 13 green; 5 green, 1 red",
			rounds: []map[string]int{
				{"green": 8, "blue": 8, "red": 20},
				{"green": 13, "blue": 8, "red": 5},
				{"green": 5, "red": 1},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			game, err := createGame(tc.description)
			if err != nil {
				t.Fatalf("error creating game: %+v\n", err)
			}

			if game.id != 3 {
				t.Errorf("game.id: want: %d, got: %d\n", 3, game.id)
			}

			if len(game.rounds) != 3 {
				t.Errorf("len(game.rounds): want: %d, got: %d\n", 3, len(game.rounds))
			}
			for i, round := range game.rounds {
				if len(round) != len(tc.rounds[i]) {
					t.Errorf("len(round): want: %d, got: %d\n", len(tc.rounds[i]), len(round))
				}

				for color, amount := range round {
					if amount != tc.rounds[i][color] {
						t.Errorf("want: %d %s, got %d %s\n", tc.rounds[i][color], color, amount, color)
					}
				}
			}
		})
	}
}

func TestSolvePartOne(t *testing.T) {
	expected := 8
	f, err := os.Open("./input_test.txt")
	if err != nil {
		t.Fatalf("error opening test input file: %+v", err)
	}

	result, err := solvePartOne(f)
	if err != nil {
		t.Fatalf("error solving day2 part1: %+v", err)
	}

	if result != expected {
		t.Errorf("want: %d, got: %d\n", expected, result)
	}
}

func TestSolvePartTwo(t *testing.T) {
	expected := 2286

	f, err := os.Open("./input_test.txt")
	if err != nil {
		t.Fatalf("error opening test input file: %+v", err)
	}

	result, err := solvePartTwo(f)
	if err != nil {
		t.Fatalf("error solving day2 part1: %+v", err)
	}

	if result != expected {
		t.Errorf("want: %d, got: %d\n", expected, result)
	}
}
