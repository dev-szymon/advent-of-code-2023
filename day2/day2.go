package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Game struct {
	id     int
	rounds []map[string]int
}

func createGame(description string) (*Game, error) {
	meta := strings.Split(description, ":")
	id, err := strconv.Atoi(strings.Fields(meta[0])[1])
	if err != nil {
		return nil, err
	}
	game := &Game{id: id, rounds: []map[string]int{}}
	rounds := strings.Split(meta[1], ";")
	for _, r := range rounds {
		round := map[string]int{}
		cubes := strings.Split(r, ", ")
		for _, c := range cubes {
			colourAndAmount := strings.Fields(c)
			colour := colourAndAmount[1]
			amount, err := strconv.Atoi(colourAndAmount[0])
			if err != nil {
				return nil, err
			}
			round[colour] += amount
		}
		game.rounds = append(game.rounds, round)
	}

	return game, nil
}

func filterValidGames(games []*Game, cubesSeed map[string]int) []*Game {
	result := []*Game{}

	for _, game := range games {
		isValid := true

	roundsLoop:
		for _, round := range game.rounds {
			for colour, amount := range round {
				if amount > cubesSeed[colour] {
					isValid = false
					break roundsLoop
				}
			}
		}

		if isValid {
			result = append(result, game)
		}
	}

	return result
}

func solvePartOne(file io.Reader) (int, error) {
	cubesSeed := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	games := []*Game{}
	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game, err := createGame(line)
		if err != nil {
			return -1, err
		}
		games = append(games, game)
	}

	validGames := filterValidGames(games, cubesSeed)
	for _, game := range validGames {
		firstRound := game.rounds[0]
		fmt.Printf("game %d: %d red, %d green, %d blue\n", game.id, firstRound["red"], firstRound["green"], firstRound["blue"])
		total += game.id
	}

	return total, nil
}

func solvePartTwo(file io.Reader) (int, error) {
	games := []*Game{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game, err := createGame(line)
		if err != nil {
			return -1, err
		}
		games = append(games, game)
	}

	gameIdTolowestPossibleAmounts := map[int]map[string]int{}
	for _, game := range games {
		gameIdTolowestPossibleAmounts[game.id] = map[string]int{}
		for _, r := range game.rounds {
			for colour, amount := range r {
				if amount > gameIdTolowestPossibleAmounts[game.id][colour] {
					gameIdTolowestPossibleAmounts[game.id][colour] = amount
				}
			}
		}
	}

	total := 0
	for _, lowestPossibleAmounts := range gameIdTolowestPossibleAmounts {
		power := 1
		for _, amount := range lowestPossibleAmounts {
			power *= amount
		}
		total += power
	}

	return total, nil
}
