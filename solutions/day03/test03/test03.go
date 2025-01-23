package main

import (
	"fmt"
	"strings"

	day03 "go-aoc-template/solutions/day03"
)

var lines = strings.Split(`101 301 501
102 302 502
103 303 503
201 401 601
202 402 602
203 403 603`, "\n")

var (
	partOneAnswer = "example answer"
	partTwoAnswer = "6"
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
	runTest(1, day03.PartOne, partOneAnswer)
	runTest(2, day03.PartTwo, partTwoAnswer)
}
