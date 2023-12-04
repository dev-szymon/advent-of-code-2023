package main

import (
	"bufio"
	"io"
	"strconv"
)

type Symbol struct {
	value       rune
	coordinates Coordinates
}

type Coordinates struct {
	y int
	x int
}

type Number struct {
	value               int
	adjacentCoordinates map[Coordinates]bool
}

var digits = map[rune]bool{
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
}
var adjacentCoordinatesOffset = map[string][]int{
	"top":         {-1, 0},
	"topRight":    {-1, 1},
	"right":       {0, 1},
	"bottomRight": {1, 1},
	"bottom":      {1, 0},
	"bottomLeft":  {1, -1},
	"left":        {0, -1},
	"topLeft":     {-1, -1},
}

func getNumbersAndSymbols(file io.Reader) ([]Symbol, []Number, error) {
	scanner := bufio.NewScanner(file)
	symbols := []Symbol{}
	numbers := []Number{}

	currentNumber := ""
	currentNumberAdjacentCoordinates := map[Coordinates]bool{}
	y := -1
	for scanner.Scan() {
		y++
		line := scanner.Text()

		for x, r := range line {
			if digits[r] {
				currentNumber += string(r)
				for _, offset := range adjacentCoordinatesOffset {
					currentNumberAdjacentCoordinates[Coordinates{y: y + offset[0], x: x + offset[1]}] = true
				}
			} else {
				if len(currentNumber) > 0 {
					n, err := strconv.Atoi(currentNumber)
					if err != nil {
						return nil, nil, err
					}
					numbers = append(numbers, Number{value: n, adjacentCoordinates: currentNumberAdjacentCoordinates})

					currentNumber = ""
					currentNumberAdjacentCoordinates = map[Coordinates]bool{}
				}

				if r == '.' {
					continue
				} else {
					symbols = append(symbols, Symbol{value: r, coordinates: Coordinates{y: y, x: x}})
				}
			}
		}
	}
	return symbols, numbers, nil
}

func solvePartOne(file io.Reader) (int, error) {
	symbols, numbers, err := getNumbersAndSymbols(file)
	if err != nil {
		return -1, err
	}

	total := 0
	for _, num := range numbers {
		for _, symbol := range symbols {
			if num.adjacentCoordinates[symbol.coordinates] {
				total += num.value
			}
		}
	}

	return total, nil
}

func solvePartTwo(file io.Reader) (int, error) {
	symbols, numbers, err := getNumbersAndSymbols(file)
	if err != nil {
		return -1, nil
	}

	total := 0
	for _, symbol := range symbols {
		adjacentNumbers := []Number{}
		for _, number := range numbers {
			if number.adjacentCoordinates[symbol.coordinates] {
				adjacentNumbers = append(adjacentNumbers, number)
			}
		}

		if len(adjacentNumbers) == 2 {
			total += adjacentNumbers[0].value * adjacentNumbers[1].value
		}
	}

	return total, nil
}
