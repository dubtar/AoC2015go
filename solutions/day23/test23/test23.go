package main

import (
	"fmt"
	"strings"

	day23 "go-aoc-template/solutions/day23"
)

var lines = strings.Split(`inc a
jio a, +2
tpl a
inc a`, "\n")

var (
	partOneAnswer = "2 0"
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
	runTest(1, day23.PartOne, partOneAnswer)
	runTest(2, day23.PartTwo, partTwoAnswer)
}
