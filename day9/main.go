package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	part1 := func() int {
		f, err := os.Open("./input.txt")
		if err != nil {
			log.Fatalf("day9: error opening input file: %+v\n", err)
		}
		defer f.Close()

		result, err := solvePartOne(f)
		if err != nil {
			log.Fatalf("day9: error solving part one: %+v\n", err)
		}

		return result
	}()

	part2 := func() int {
		f, err := os.Open("./input.txt")
		if err != nil {
			log.Fatalf("day9: error opening input file: %+v\n", err)
		}
		defer f.Close()

		result, err := solvePartTwo(f)
		if err != nil {
			log.Fatalf("day9: error opening input file: %+v\n", err)
		}
		return result
	}()

	fmt.Printf("day9: part1: %d\n", part1)
	fmt.Printf("day9: part2: %d\n", part2)
}
