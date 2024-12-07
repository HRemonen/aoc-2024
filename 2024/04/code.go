package main

import (
	"fmt"
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

// The possible variations are:
// S.S
// .A.
// M.M

// S.M
// .A.
// S.M

// M.M
// .A.
// S.S

// M.S
// .A.
// M.S
func solve2(input []string) int {
	result := 0

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {

			if input[row][col] == 'A' {
				if row > 0 && row < len(input)-1 && col > 0 && col < len(input[row])-1 {
					topLeft := input[row-1][col-1]
					topRight := input[row-1][col+1]
					bottomLeft := input[row+1][col-1]
					bottomRight := input[row+1][col+1]

					if (topLeft == 'S' && topRight == 'S' && bottomLeft == 'M' && bottomRight == 'M') || // Pattern 1
						(topLeft == 'S' && topRight == 'M' && bottomLeft == 'S' && bottomRight == 'M') || // Pattern 2
						(topLeft == 'M' && topRight == 'M' && bottomLeft == 'S' && bottomRight == 'S') || // Pattern 3
						(topLeft == 'M' && topRight == 'S' && bottomLeft == 'M' && bottomRight == 'S') { // Pattern 4
						result++
						fmt.Printf("X-MAS pattern found at (%d, %d): TL=%c, TR=%c, BL=%c, BR=%c\n",
							row, col, topLeft, topRight, bottomLeft, bottomRight)
					}
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
		return solve2(rows)
	}

	return solve(rows)
}
