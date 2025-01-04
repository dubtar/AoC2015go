package main

import (
	"fmt"
	"strings"

	day22 "go-aoc-template/solutions/day22"
)

var lines = strings.Split(`example input`, "\n")

var (
	partOneAnswer = ""
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
	fmt.Printf("%v", day22.Run(day22.TestGame()))
	// runTest(1, day22.PartOne, partOneAnswer)
	// runTest(2, day22.PartTwo, partTwoAnswer)
}
