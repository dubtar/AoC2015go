package main

import (
	"fmt"
	"strings"

	day23 "go-aoc-template/solutions/day23"
)

var lines_t = strings.Split(`cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a`, "\n")

var lines = strings.Split(`
cpy a b
dec b
cpy a d
cpy 0 a
cpy b c
add a c # inc a
# dec c
# jnz c -2
dec d
jnz d -5
dec b
cpy b c
cpy c d
# dec d
add c d # inc c
# jnz d -2
tgl c
cpy -16 c
jnz 1 c
cpy 71 c
jnz 72 d
add a d #inc a
# inc d
# jnz d -2
inc c
jnz c -5`, "\n")

var (
	partOneAnswer = "3"
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
	runTest(1, day23.PartOne, partOneAnswer)
	runTest(2, day23.PartTwo, partTwoAnswer)
}
