package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func isIncreasing(input []string) bool {
	for i := 1; i < len(input); i++ {
		prev, _ := strconv.Atoi(string(input[i-1]))
		curr, _ := strconv.Atoi(string(input[i]))
		if prev >= curr {
			return false
		}
	}
	return true
}

// Check if the levels are strictly decreasing
func isDecreasing(input []string) bool {
	for i := 1; i < len(input); i++ {
		prev, _ := strconv.Atoi(string(input[i-1]))
		curr, _ := strconv.Atoi(string(input[i]))
		if prev <= curr {
			return false
		}
	}
	return true
}

// The Red-Nosed reactor safety systems can only tolerate levels that are
// either gradually increasing or gradually decreasing.
// So, a report only counts as safe if both of the following are true:
//
// - The levels are either all increasing or all decreasing.
// - Any two adjacent levels differ by at least one and at most three.
func checkRules(input []string) bool {
	if !isIncreasing(input) && !isDecreasing(input) {
		return false
	}

	for i := 1; i < len(input); i++ {
		prev, _ := strconv.Atoi(string(input[i-1]))
		curr, _ := strconv.Atoi(string(input[i]))

		diff := math.Abs(float64(prev - curr))
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	result := 0

	for _, line := range strings.Split(input, "\n") {
		if checkRules(strings.Fields(line)) {
			result++
		}
	}
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return result
}
