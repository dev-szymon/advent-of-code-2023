package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers map[int]bool
	playerNumbers  []int
}

func getCards(file io.Reader) ([]*Card, error) {
	scanner := bufio.NewScanner(file)

	cards := []*Card{}
	for scanner.Scan() {
		description := strings.Split(scanner.Text(), ":")
		card := &Card{winningNumbers: map[int]bool{}, playerNumbers: []int{}}
		_, err := fmt.Sscanf(description[0], "Card %d", &card.id)
		if err != nil {
			return nil, err
		}

		numbers := strings.Split(description[1], "|")
		winningNumbers := numbers[0]
		for _, number := range strings.Fields(winningNumbers) {
			n, err := strconv.Atoi(number)
			if err != nil {
				return nil, err
			}
			card.winningNumbers[n] = true
		}

		playerNumbers := numbers[1]
		for _, number := range strings.Fields(playerNumbers) {
			n, err := strconv.Atoi(number)
			if err != nil {
				return nil, err
			}
			card.playerNumbers = append(card.playerNumbers, n)
		}

		cards = append(cards, card)
	}

	return cards, nil
}

func solvePartOne(file io.Reader) (int, error) {
	cards, err := getCards(file)
	if err != nil {
		return -1, err
	}

	total := 0
	for _, card := range cards {
		score := 0
		for _, playerNumber := range card.playerNumbers {
			if card.winningNumbers[playerNumber] {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		total += score
	}

	return total, nil
}

func solvePartTwo(file io.Reader) (int, error) {
	cards, err := getCards(file)
	if err != nil {
		return -1, err
	}

	cardIdToCopies := map[int]int{}
	for _, card := range cards {
		cardIdToCopies[card.id] = 1
	}

	for _, card := range cards {
		score := 0
		for _, playerNumber := range card.playerNumbers {
			if card.winningNumbers[playerNumber] {
				score++
			}
		}

		for i := 1; i < score+1; i++ {
			cardIdToCopies[card.id+i] += cardIdToCopies[card.id]
		}
	}

	total := 0
	for _, copies := range cardIdToCopies {
		total += copies
	}

	return total, nil
}
