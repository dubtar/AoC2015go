package main

import (
	"fmt"
	"strings"

	day07 "go-aoc-template/solutions/day07"
)

var lines = strings.Split(`abba[mnop]qrst
abcd[bddb]xyyx
abaa[babr]tyui
ioxxoj[asdfgh]zxcvbn`, "\n")

var (
	partOneAnswer = "2"
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
	runTest(1, day07.PartOne, partOneAnswer)
	runTest(2, day07.PartTwo, partTwoAnswer)
}
