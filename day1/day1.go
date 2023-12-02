package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

var digitReplacements = map[string]string{
	"1": "1",
	"2": "2",
	"3": "3",
	"4": "4",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",
}
var wordReplacements = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func findCalibrationValue(s string, withWords bool) (int, error) {
	replacements := digitReplacements
	var (
		result = ""
		copy   = s[:]
		lookup = ""
	)

	if withWords {
		for k, v := range wordReplacements {
			replacements[k] = v
		}
	}

ltr:
	for len(copy) > 0 {
		lookup += string(copy[0])
		copy = copy[1:]

		for old, new := range replacements {
			if strings.HasSuffix(lookup, old) {
				result += new
				break ltr
			}
		}
	}

	lookup = ""
	copy = s[:]

rtl:
	for len(copy) > 0 {
		lookup = string(copy[len(copy)-1]) + lookup
		copy = copy[:len(copy)-1]

		for old, new := range replacements {
			if strings.HasPrefix(lookup, old) {
				result += new
				break rtl
			}
		}
	}

	n, err := strconv.Atoi(result)
	if err != nil {
		return -1, err
	}
	return n, err
}

func solvePartOne(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		n, err := findCalibrationValue(line, false)
		if err != nil {
			return -1, err
		}

		total += n
	}

	return total, nil
}

func solvePartTwo(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		n, err := findCalibrationValue(line, true)
		if err != nil {
			return -1, err
		}

		total += n
	}

	return total, nil
}
