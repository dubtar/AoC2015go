package day23

import (
	"fmt"
	day12 "go-aoc-template/solutions/day12"
)

func PartOne(lines []string) string {
	computer := day12.NewComputer(lines)
	computer.Set("a", 7)
	computer.Run()
	return fmt.Sprint(computer.Get("a"))
}
func PartTwo(lines []string) string {
	computer := day12.NewComputer(lines)
	computer.Set("a", 12)
	computer.Run()
	return fmt.Sprint(computer.Get("a"))
}
