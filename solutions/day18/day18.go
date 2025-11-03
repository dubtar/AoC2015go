package day18

import (
	"fmt"
	"strconv"
)

// import (
// 	h "go-aoc-template/internal/helpers"
// )

func isTrap(prevline []int, index int) int {
	left := index > 0 && prevline[index-1] == 1
	center := prevline[index] == 1
	right := index < len(prevline)-1 && prevline[index+1] == 1
	if left && center && !right ||
		!left && center && right ||
		left && !center && !right ||
		!left && !center && right {
		return 1
	} else {
		return 0
	}
}
func printLine(line []int) {
	for _, c := range line {
		if c == 1 {
			fmt.Print("^")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func PartOne(lines []string) string {
	steps := 40
	if len(lines[0]) < 30 {
		steps = 10
	}
	return run(lines, steps)
}
func run(lines []string, steps int) string {
	var result int = 0
	line := make([]int, len(lines[0]))
	for i, c := range lines[0] {
		if c == '^' {
			line[i] = 1
		} else {
			line[i] = 0
			result += 1
		}
	}
	// printLine(line)
	for _ = range steps - 1 {
		nextLine := make([]int, len(line))
		for i := range line {
			nextLine[i] = isTrap(line, i)
			if nextLine[i] == 0 {
				result += 1
			}
		}
		line = nextLine
		// printLine(line)
	}
	return strconv.Itoa(result)
}

func PartTwo(lines []string) string {
	return run(lines, 400000)
}
