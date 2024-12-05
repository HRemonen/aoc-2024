package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	array1 := []string{}
	array2 := []string{}

	for i, v := range strings.Fields(input) {
		if i%2 == 0 {
			array1 = append(array1, v)
		} else {
			array2 = append(array2, v)
		}
	}

	// sort both arrays
	sort.Strings(array1)
	sort.Strings(array2)

	var result int
	for i := 0; i < len(array1); i++ {
		val1, _ := strconv.Atoi(array1[i])
		val2, _ := strconv.Atoi(array2[i])

		result += diff(val1, val2)
	}

	if part2 {
		var result2 int
		// find out how many time an element from the array1 is in array2
		for _, v := range array1 {
			times := 0
			for _, w := range array2 {
				if v == w {
					times++
				}
			}
			val, _ := strconv.Atoi(v)
			result2 += val * times
		}
		return result2
	}
	return result
}
