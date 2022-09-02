package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/imandradesign/aoc2017.git/pkg/helpers"
)

func parseLine(line string) []int {
	// "line": []string{"5 1 9 5", "7 5 3", "2 4 6 8"}
	theInts := []int{}

	for _, input := range strings.Fields(line) {
		theIntProbably, err := strconv.Atoi(input)
		if err != nil {
			log.Fatalf("Failed to make an int out of %s: %+v", input, err)
		}

		theInts = append(theInts, theIntProbably)
	}

	return theInts
}

func parseData(data []string) [][]int {
	theOutput := [][]int{}

	for _, line := range data {
		if len(strings.TrimSpace(line)) > 0 {
			ints := parseLine(line)

			theOutput = append(theOutput, ints)
		}
	}

	return theOutput
}

func sortOurShitOut(theInts []int) (int, int) {
	sort.Ints(theInts)
	return theInts[0], theInts[len(theInts)-1]
}

func modulusOurShitOut(data []int) (int, int) {
	sort.Ints(data)

	var foundNumerator = 0
	var foundDenominator = 0

	for idx, denominator := range data {
		for _, numerator := range data[idx+1:] {
			if numerator%denominator == 0 {
				foundNumerator = numerator
				foundDenominator = denominator
			}
		}
	}

	return foundNumerator, foundDenominator
}

func solve(parsed [][]int, numberCruncher func([]int) (int, int), resulter func(int, int) int) int {

	theResult := 0

	for _, row := range parsed {
		one, two := numberCruncher(row)
		result := resulter(one, two)

		theResult += result
	}

	return theResult
}

// Day 2: Corruption Checksum

// For each row, determine the difference between the largest value and the smallest value; the checksum is the sum of all of these differences.

// For example, given the following spreadsheet:

// 5 1 9 5
// 7 5 3
// 2 4 6 8

//     The first row's largest and smallest values are 9 and 1, and their difference is 8.
//     The second row's largest and smallest values are 7 and 3, and their difference is 4.
//     The third row's difference is 6.

// In this example, the spreadsheet's checksum would be 8 + 4 + 6 = 18.

func partOne(parsed [][]int) int {
	return solve(parsed, sortOurShitOut, func(min int, max int) int { return max - min })
}

// Part Two!
// It sounds like the goal is to find the only two numbers in each row where one evenly divides the other - that is, where the result of the division operation is a whole number. They would like you to find those numbers on each line, divide them, and add up each line's result.

// For example, given the following spreadsheet:

// 5 9 2 8
// 9 4 7 3
// 3 8 6 5

//     In the first row, the only two numbers that evenly divide are 8 and 2; the result of this division is 4.
//     In the second row, the two numbers are 9 and 3; the result is 3.
//     In the third row, the result is 2.

// In this example, the sum of the results would be 4 + 3 + 2 = 9.

func partTwo(parsed [][]int) int {
	return solve(parsed, modulusOurShitOut, func(numerator int, denominator int) int { return numerator / denominator })
}

func main() {
	data, err := helpers.LoadInputData("2")

	if err != nil {
		log.Fatalf("Oh shit, couldn't load data: %+v", err)
	}

	parsed := parseData(data)
	fmt.Println("Day Two!")
	fmt.Printf("Part One => %d\n", partOne(parsed))
	fmt.Printf("Part Two => %d\n", partTwo(parsed))
}
