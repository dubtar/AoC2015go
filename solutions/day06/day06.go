package day06

// import (
//     h "go-aoc-template/internal/helpers"
// )

func PartOne(lines []string) string {
	size := len(lines[0])
	maps := make([]map[rune]int, size)
	best := make([]struct {
		letter rune
		count  int
	}, size)
	for i := range maps {
		maps[i] = make(map[rune]int)
	}
	for _, line := range lines {
		for i, c := range line {
			maps[i][c]++
			count := maps[i][c]
			if count > best[i].count {
				best[i].letter = c
				best[i].count = count
			}
		}
	}
	result := ""
	for _, b := range best {
		result += string(b.letter)
	}
	return result
}

func PartTwo(lines []string) string {
	size := len(lines[0])
	counts := make([]map[rune]int, size)
	for i := range counts {
		counts[i] = make(map[rune]int)
	}
	for _, line := range lines {
		for i, c := range line {
			counts[i][c]++
		}
	}
	result := ""
	for _, m := range counts {
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
