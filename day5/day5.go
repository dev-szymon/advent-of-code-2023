package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type Mapping struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func parseAlmanac(file io.Reader) ([]int, [][]Mapping, error) {
	scanner := bufio.NewScanner(file)

	seeds := []int{}
	scanner.Scan()
	for _, s := range strings.Fields(strings.Split(scanner.Text(), ":")[1]) {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, nil, err
		}
		seeds = append(seeds, n)
	}

	mappings := [][]Mapping{}
	currentStep := -1
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if line[len(line)-1] == ':' {
			currentStep++
			mappings = append(mappings, []Mapping{})
			continue
		}

		mapping := Mapping{}
		_, err := fmt.Sscanf(line, "%d %d %d", &mapping.destinationRangeStart, &mapping.sourceRangeStart, &mapping.rangeLength)
		if err != nil {
			return nil, nil, err
		}

		mappings[currentStep] = append(mappings[currentStep], mapping)
	}

	return seeds, mappings, nil
}

func getDestinationValueFromSource(source int, mappings []Mapping) int {
	for _, m := range mappings {
		if source >= m.sourceRangeStart && source < m.sourceRangeStart+m.rangeLength {
			return source + (m.destinationRangeStart - m.sourceRangeStart)
		}
	}

	return source
}

func solvePartOne(file io.Reader) (int, error) {
	seeds, mappings, err := parseAlmanac(file)
	if err != nil {
		return -1, err
	}
	lowestLocation := math.MaxInt
	for _, s := range seeds {
		location := s
		for _, m := range mappings {
			location = getDestinationValueFromSource(location, m)
		}

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation, nil
}

func solvePartTwo(file io.Reader) (int, error) {
	seeds, mappings, err := parseAlmanac(file)
	if err != nil {
		return -1, err
	}

	seedStartToRange := map[int]int{}
	for i, r := range seeds {
		if i > 0 && i%2 == 1 {
			seedStartToRange[seeds[i-1]] = r
		}
	}

	lowestLocation := math.MaxInt
	for seedStart, seedRange := range seedStartToRange {
		for i := seedStart; i < seedStart+seedRange; i++ {
			nextLocation := i
			for _, m := range mappings {
				nextLocation = getDestinationValueFromSource(nextLocation, m)
			}

			if nextLocation < lowestLocation {
				lowestLocation = nextLocation
			}
		}
	}

	return lowestLocation, nil
}
