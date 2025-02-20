package main

import (
	"fmt"
	"strings"

	day20 "go-aoc-template/solutions/day20"
)

var lines = strings.Split(`30470`, "\n")

var (
	partOneAnswer = "960" // 480, 240, 120, 60, 30, 15, 3, 5
	// 98280 /2 = 49140 /2 = 24570 /2 = 12285 /3 = 4095/3=1365/3=455/5=91
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
	runTest(1, day20.PartOne, partOneAnswer)
	runTest(2, day20.PartTwo, partTwoAnswer)
}
