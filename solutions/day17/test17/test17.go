package main

import (
	"fmt"
	"strconv"
	"strings"

	day17 "go-aoc-template/solutions/day17"
)

var lines = strings.Split(`20
15
10
5
5`, "\n")

var (
	partOneAnswer = "4 3"
	partTwoAnswer = "3"
)

type SolutionFunc func([]string, int64) (int, int)

func runTest(part int, solution SolutionFunc, expected string) {
	fmt.Printf("Part %d: ", part)
	result1, result2 := solution(lines, 25)
	result := strconv.Itoa(result1) + " " + strconv.Itoa(result2)
	if result != expected {
		fmt.Printf("\033[31m%v\033[0m (expected \033[32m%v\033[0m)\n", result, expected)
	} else {
		fmt.Printf("\033[32m%v\033[0m\n", result)
	}
}

func main() {
	runTest(1, day17.PartOneTargeted, partOneAnswer)
	// runTest(2, day17.PartOneTargeted, partTwoAnswer)
}
