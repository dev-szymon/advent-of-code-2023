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
			log.Fatalf("day8: error opening input file: %+v\n", err)
		}
		defer f.Close()

		return solvePartOne(f)
	}()

	part2 := func() int {
		f, err := os.Open("./input.txt")
		if err != nil {
			log.Fatalf("day8: error opening input file: %+v\n", err)
		}
		defer f.Close()

		return solvePartTwo(f)
	}()

	fmt.Printf("day8: part1: %d\n", part1)
	fmt.Printf("day8: part2: %d\n", part2)
}
