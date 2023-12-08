package main

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
)

var baseCardToScore = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func replaceJacksWithJokers(cToS map[rune]int) map[rune]int {
	var cardToScoreWithJokers = map[rune]int{}
	for card, score := range cToS {
		if card == 'J' {
			cardToScoreWithJokers[card] = 1
		} else {
			cardToScoreWithJokers[card] = score
		}
	}
	return cardToScoreWithJokers
}

type Hand struct {
	cards   string
	bid     int
	figures map[rune]int
}

func newHand(cards string, bid int) Hand {
	hand := Hand{
		cards:   cards,
		bid:     bid,
		figures: map[rune]int{},
	}

	for _, card := range cards {
		hand.figures[card]++
	}

	return hand
}

const (
	typeHighCard = iota
	typeOnePair
	typeTwoPair
	typeThreeOfKind
	typeFullHouse
	typeFourOfKind
	typeFiveOfKind
)

func (h Hand) getHandType() int {
	highestCount := 0
	for _, count := range h.figures {
		if count > highestCount {
			highestCount = count
		}
	}

	switch highestCount {
	case 2:
		if len(h.figures) == 3 {
			return typeTwoPair
		} else {
			return typeOnePair
		}
	case 3:
		if len(h.figures) == 2 {
			return typeFullHouse
		} else {
			return typeThreeOfKind
		}
	case 4:
		return typeFourOfKind
	case 5:
		return typeFiveOfKind
	default:
		return typeHighCard
	}
}

func (h Hand) getHandTypeWithJokers() int {
	jokers := h.figures['J']

	highestCount := 0
	for card, count := range h.figures {
		if card != 'J' && count > highestCount {
			highestCount = count
		}
	}

	totalHighest := jokers + highestCount

	switch totalHighest {
	case 2:
		pairs := 0
		for _, count := range h.figures {
			if count == 2 {
				pairs++
			}
		}
		if pairs == 2 || (pairs == 1 && jokers == 1) {
			return typeTwoPair
		} else {
			return typeOnePair
		}
	case 3:
		if (jokers > 0 && len(h.figures) == 3) || len(h.figures) == 2 {
			return typeFullHouse
		} else {
			return typeThreeOfKind
		}
	case 4:
		return typeFourOfKind
	case 5:
		return typeFiveOfKind
	default:
		return typeHighCard
	}
}

func compareHands(a, b Hand, withJokers bool) int {
	var aType int
	var bType int
	var cardToScore map[rune]int
	if withJokers {
		aType = a.getHandTypeWithJokers()
		bType = b.getHandTypeWithJokers()
		cardToScore = replaceJacksWithJokers(baseCardToScore)
	} else {
		aType = a.getHandType()
		bType = b.getHandType()
		cardToScore = baseCardToScore
	}

	if aType < bType {
		return -1
	} else if aType > bType {
		return 1
	} else if a.cards == b.cards {
		return 0
	} else {
		for i, c := range a.cards {
			if cardToScore[c] == cardToScore[rune(b.cards[i])] {
				continue
			} else if cardToScore[c] < cardToScore[rune(b.cards[i])] {
				return -1
			} else {
				return 1
			}
		}
		return 0
	}
}

func solvePartOne(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)

	hands := []Hand{}
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		cards := line[0]
		bid, err := strconv.Atoi(line[1])
		if err != nil {
			return -1, err
		}
		hands = append(hands, newHand(cards, bid))
	}

	slices.SortFunc(hands, func(a, b Hand) int { return compareHands(a, b, false) })

	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}
	return total, nil
}

func solvePartTwo(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)

	hands := []Hand{}
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		cards := line[0]
		bid, err := strconv.Atoi(line[1])
		if err != nil {
			return -1, err
		}
		hands = append(hands, newHand(cards, bid))
	}

	slices.SortFunc(hands, func(a, b Hand) int { return compareHands(a, b, true) })

	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}
	return total, nil
}
