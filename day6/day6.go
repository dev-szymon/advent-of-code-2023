package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func canBeatCurrentRecord(speed, time, currentRecord int) bool {
	distance := speed * (time - speed)
	return distance > currentRecord
}

func getWinningScenariosCount(time, distance int) int {
	low := time
	for {
		low = low / 2
		if !canBeatCurrentRecord(low, time, distance) {
			break
		}
	}
	for {
		low++
		if canBeatCurrentRecord(low, time, distance) {
			break
		}
	}

	high := 0
	for {
		high = (time - high/2)
		if !canBeatCurrentRecord(high, time, distance) {
			break
		}
	}
	for {
		high--
		if canBeatCurrentRecord(high, time, distance) {
			break
		}
	}
	return high - low + 1
}

func solvePartOne(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	times := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	scanner.Scan()
	distances := strings.Fields(strings.Split(scanner.Text(), ":")[1])

	timeToDistance := map[int]int{}
	for i := 0; i < len(times); i++ {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			return -1, err
		}

		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			return -1, err
		}

		timeToDistance[time] = distance
	}

	total := 1
	for time, distance := range timeToDistance {
		total *= getWinningScenariosCount(time, distance)
	}

	return total, nil
}

func solvePartTwo(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	time, err := strconv.Atoi(strings.Join(strings.Fields(strings.Split(scanner.Text(), ":")[1]), ""))
	if err != nil {
		return -1, err
	}
	scanner.Scan()
	distance, err := strconv.Atoi(strings.Join(strings.Fields(strings.Split(scanner.Text(), ":")[1]), ""))
	if err != nil {
		return -1, err
	}

	total := getWinningScenariosCount(time, distance)

	return total, nil
}
