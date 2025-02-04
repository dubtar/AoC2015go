package main

import (
	"fmt"
	"strings"

	day08 "go-aoc-template/solutions/day08"
)

var lines = strings.Split(`rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4
rotate column x=1 by 1`, "\n")

var (
	partOneAnswer = "example answer"
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
	runTest(1, day08.PartOne, partOneAnswer)
	runTest(2, day08.PartTwo, partTwoAnswer)
}
