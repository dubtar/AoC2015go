package main

import (
	"fmt"
	"strings"

	day01 "go-aoc-template/solutions/day01"
)

var lines = strings.Split(`)())())`, "\n")

var (
	partOneAnswer = "-3"
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
	runTest(1, day01.PartOne, partOneAnswer)
	runTest(2, day01.PartTwo, partTwoAnswer)
}
