package main

import (
	"strconv"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	result := 0
	tokenBuilder := ""
	leftOperand := ""
	leftSeen := false
	rightOperand := ""
	rightSeen := false

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case 'm':
			if len(tokenBuilder) == 0 {
				tokenBuilder += "m"
			}
		case 'u':
			if tokenBuilder == "m" {
				tokenBuilder += "u"
			}
		case 'l':
			if tokenBuilder == "mu" {
				tokenBuilder += "l"
			}
		case '(':
			if tokenBuilder == "mul" {
				tokenBuilder += "("

				// scan numbers until we hit a ','
				scannedLen := 0
				for i++; i < len(input); i++ {
					if scannedLen == 4 {
						// revert head to the start if too many characters are scanned
						i -= 3
						break
					}
					if !leftSeen && input[i] == ',' {
						leftSeen = true
						scannedLen = 0
						i++
					}
					if leftSeen && !rightSeen && input[i] == ')' {
						rightSeen = true
						scannedLen = 0
						i--
						break
					}

					if !leftSeen {
						leftOperand += string(input[i])
					} else if !rightSeen {
						rightOperand += string(input[i])
					}
					scannedLen++
				}

			}
		case ')':
			if tokenBuilder == "mul(" && leftSeen && rightSeen {
				// commit the operation to result
				leftInt, _ := strconv.Atoi(leftOperand)
				rightInt, _ := strconv.Atoi(rightOperand)
				result += leftInt * rightInt

				// reset
				tokenBuilder = ""
				leftOperand = ""
				leftSeen = false
				rightOperand = ""
				rightSeen = false
			}
		default:
			tokenBuilder = ""
			leftOperand = ""
			leftSeen = false
			rightOperand = ""
			rightSeen = false
		}
	}
	return result
}
