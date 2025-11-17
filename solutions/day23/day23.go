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
	if len(lines) > 20 && lines[5] == "inc a" {
		// Специфичная под задание оптимизация
		lines[5] = "addmul a c d # inc a"
		lines[6] = "cpy 0 c #dec c"
		lines[7] = "# jnz c -2"
		lines[8] = "cpy 0 d #dec d"
	}
	computer := day12.NewComputer(lines)
	computer.Set("a", 12)
	computer.Run()
	return fmt.Sprint(computer.Get("a"))
}
