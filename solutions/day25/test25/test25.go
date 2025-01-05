package main

import (
	"fmt"
	"strings"

	day25 "go-aoc-template/solutions/day25"
)

var lines = strings.Split(`To continue, please consult the code grid in the manual.  Enter the code at row 6, column 6.`, "\n")

var (
	partOneAnswer = "27995004"
	partTwoAnswer = "example answer"
)

type SolutionFunc func([]string) string

func runTest(part int, solution SolutionFunc, expected string) {
	fmt.Printf("Part %d: ", part)
	result := solution(lines)
	if result != expected {
		fmt.Printf("\033[31m%v\033[0m (expected \033[32m%v\033[0m)\n", result, expected)
	} else {
		fmt.Printf("\033[32m%v\033[0m\n", result)
	}
}

func main() {
	runTest(1, day25.PartOne, partOneAnswer)
	runTest(2, day25.PartTwo, partTwoAnswer)
}
