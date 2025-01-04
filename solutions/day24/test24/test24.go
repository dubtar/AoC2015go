package main

import (
	"fmt"
	"strings"

	day24 "go-aoc-template/solutions/day24"
)

var lines = strings.Split(`1
2
3
4
5
7
8
9
10
11`, "\n")

var (
	partOneAnswer = "99"
	partTwoAnswer = "44"
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
	runTest(1, day24.PartOne, partOneAnswer)
	runTest(2, day24.PartTwo, partTwoAnswer)
}
