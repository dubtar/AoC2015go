package main

import (
	"fmt"
	"strings"

	day23 "go-aoc-template/solutions/day23"
)

var lines = strings.Split(`cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a`, "\n")

var (
	partOneAnswer = "3"
	partTwoAnswer = "3"
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
