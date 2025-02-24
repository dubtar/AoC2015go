package main

import (
	"fmt"
	"strings"

	day15 "go-aoc-template/solutions/day15"
)

var lines = strings.Split(`Disc #1 has 5 positions; at time=0, it is at position 4.
Disc #2 has 2 positions; at time=0, it is at position 1.`, "\n")

var (
	partOneAnswer = "5"
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
	runTest(1, day15.PartOne, partOneAnswer)
	runTest(2, day15.PartTwo, partTwoAnswer)
}
