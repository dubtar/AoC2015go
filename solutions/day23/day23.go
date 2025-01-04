package day23

import (
	"fmt"
	h "go-aoc-template/internal/helpers"
	"strings"
)

func Run(a uint64, lines []string) string {
	var b uint64
	for i := 0; i >= 0 && i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		reg := &b
		if parts[1][0] == 'a' {
			reg = &a
		}
		switch parts[0] {
		case "hlf":
			*reg /= 2
		case "tpl":
			*reg *= 3
		case "inc":
			*reg++
		case "jmp":
			i += int(h.ToInt(parts[1]) - 1)
		case "jie":
			if *reg%2 == 0 {
				i += int(h.ToInt(parts[2]) - 1)
			}
		case "jio":
			if *reg == 1 {
				i += int(h.ToInt(parts[2]) - 1)
			}
		}
	}
	return fmt.Sprint(a, b)
}

func PartOne(lines []string) string {
	return Run(0, lines)
}
func PartTwo(lines []string) string {
	return Run(1, lines)
}
