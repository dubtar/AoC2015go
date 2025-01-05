package day01

import "fmt"

func PartOne(lines []string) string {
	floor := 0
	for _, c := range lines[0] {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
	}
	return fmt.Sprint(floor)
}

func PartTwo(lines []string) string {
	floor := 0
	for i, c := range lines[0] {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
		if floor == -1 {
			return fmt.Sprint(i + 1)
		}
	}
	panic("")
}
