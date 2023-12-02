package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	part1 := func() int {
		f1, err := os.Open("./input.txt")
		if err != nil {
			log.Fatalf("Error opening input file for part1: %+v", err)
		}
		defer f1.Close()

		result, err := solvePartOne(f1)
		if err != nil {
			log.Fatalf("Error solving part1: %+v", err)
		}
		return result
	}()

	part2 := func() int {
		f2, err := os.Open("./input.txt")
		if err != nil {
			log.Fatalf("Error opening input file for part1: %+v", err)
		}
		defer f2.Close()

		result, err := solvePartTwo(f2)
		if err != nil {
			log.Fatalf("Error solving part2: %+v", err)
		}
		return result
	}()

	fmt.Printf("day2 part1: %d\n", part1)
	fmt.Printf("day2 part2: %d\n", part2)
}
