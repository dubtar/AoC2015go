package main

import (
	"fmt"
	"strings"

	day19 "go-aoc-template/solutions/day19"
)

var lines = strings.Split(`5`, "\n") // HOH

var (
	partOneAnswer = "3"
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
	runTest(1, day19.PartOne, partOneAnswer)
	runTest(2, day19.PartTwo, partTwoAnswer)
}
