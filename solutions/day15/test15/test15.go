package main

import (
	"fmt"
	"strings"

	day15 "go-aoc-template/solutions/day15"
)

var lines = strings.Split(`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`, "\n")

var (
	partOneAnswer = "62842880"
	partTwoAnswer = "57600000"
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
