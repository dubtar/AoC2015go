package main

import (
	"fmt"
	"slices"
	"strings"

	day21 "go-aoc-template/solutions/day21"
)

var lines = strings.Split(`swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d
`, "\n")

var (
	partOneAnswer = "decab"
	partTwoAnswer = "abcde"
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
func runCheck() {
	const line = "abcdefgh"
	for i := range len(line) {
		value := strings.Split(line, "")
		target := value[i]

		// forward
		ind := slices.Index(value, target)
		if ind >= 4 {
			ind += 1
		}
		ind++
		ind = len(value) - (ind % len(value))
		value = append(value[ind:], value[:ind]...)
		fmt.Printf("%d: %s ", i, strings.Join(value, ""))

		//back
		var rotates = []int{7, 0, 4, 1, 5, 2, 6, 3}
		ind = slices.Index(value, target)
		ind = (len(value) + ind - rotates[ind]) % len(value) // reverse
		value = append(value[ind:], value[:ind]...)
		fmt.Println(strings.Join(value, "")) // should be
	}
}
func main() {
	// runCheck()
	runTest(1, day21.PartOne, partOneAnswer)
	runTest(2, day21.PartTwo, partTwoAnswer)
}
