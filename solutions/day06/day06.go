package day06

// import (
//     h "go-aoc-template/internal/helpers"
// )

func PartOne(lines []string) string {
	size := len(lines[0])
	maps := make([]map[rune]int, size)
	for i := range maps {
		maps[i] = make(map[rune]int)
	}
	for _, line := range lines {
		for i, c := range line {
			maps[i][c]++
		}
	}
	result := ""
	for _, m := range maps {
		best := ' '
		bestCount := 0
		for c, count := range m {
			if count > bestCount {
				best = c
				bestCount = count
			}
		}
		result += string(best)
	}
	return result
}

func PartTwo(lines []string) string {
	size := len(lines[0])
	maps := make([]map[rune]int, size)
	for i := range maps {
		maps[i] = make(map[rune]int)
	}
	for _, line := range lines {
		for i, c := range line {
			maps[i][c]++
		}
	}
	result := ""
	for _, m := range maps {
		best := ' '
		bestCount := len(lines)
		for c, count := range m {
			if count < bestCount {
				best = c
				bestCount = count
			}
		}
		result += string(best)
	}
	return result
}
