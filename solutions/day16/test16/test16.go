package main

import (
	"fmt"
	"strings"

	day16 "go-aoc-template/solutions/day16"
)

var lines = strings.Split(`Sue 1: children: 3, goldfish: 5, vizslas: 0
Sue 2: perfumes: 1, trees: 6, goldfish: 0`, "\n")

var (
	partOneAnswer = "1"
	partTwoAnswer = "2"
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
	runTest(1, day16.PartOne, partOneAnswer)
	runTest(2, day16.PartTwo, partTwoAnswer)
}
