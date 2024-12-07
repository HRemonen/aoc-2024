package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

var directions = [][2]int{
	{-1, 0},  // up
	{1, 0},   // down
	{0, -1},  // left
	{0, 1},   // right
	{-1, -1}, // top-left diagonal
	{-1, 1},  // top-right diagonal
	{1, -1},  // bottom-left diagonal
	{1, 1},   // bottom-right diagonal
}

func solve(input []string) int {
	word := "XMAS"
	wordLength := len(word)
	result := 0

	for row := range input {
		for col := range input[row] {
			for _, dir := range directions {
				match := true
				for i := 0; i < wordLength; i++ {
					newRow := row + i*dir[0]
					newCol := col + i*dir[1]

					if newRow < 0 || newRow >= len(input) || newCol < 0 || newCol >= len(input[row]) {
						match = false
						break
					}
					if input[newRow][newCol] != word[i] {
						match = false
						break
					}
				}
				if match {
					result++
				}
			}
		}
	}

	return result
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	rows := strings.Split(input, "\n")

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	result1 := solve(rows)

	return result1
}
