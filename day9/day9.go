package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

const (
	forward = iota
	backward
)

func extrapolateValue(values []int, direction int) int {
	hasReachedZeroPoint := true
	for _, n := range values {
		if n != 0 {
			hasReachedZeroPoint = false
			break
		}
	}
	if hasReachedZeroPoint {
		return 0
	}

	nextRow := []int{}
	for i, item := range values {
		if i == 0 {
			continue
		} else {
			nextRow = append(nextRow, item-values[i-1])
		}
	}

	if direction == forward {
		return values[len(values)-1] + extrapolateValue(nextRow, direction)

	} else {
		return values[0] - extrapolateValue(nextRow, direction)
	}
}

func solvePartOne(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)

	lines := [][]int{}
	for scanner.Scan() {
		lineValues := []int{}
		for _, value := range strings.Fields(scanner.Text()) {
			n, err := strconv.Atoi(value)
			if err != nil {
				return -1, err
			}
			lineValues = append(lineValues, n)
		}
		lines = append(lines, lineValues)
	}

	total := 0
	for _, line := range lines {
		total += extrapolateValue(line, forward)
	}

	return total, nil
}

func solvePartTwo(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)

	lines := [][]int{}
	for scanner.Scan() {
		lineValues := []int{}
		for _, value := range strings.Fields(scanner.Text()) {
			n, err := strconv.Atoi(value)
			if err != nil {
				return -1, err
			}
			lineValues = append(lineValues, n)
		}
		lines = append(lines, lineValues)
	}

	total := 0
	for _, line := range lines {
		total += extrapolateValue(line, backward)
	}

	return total, nil
}
