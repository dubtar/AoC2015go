package day04

import (
	h "go-aoc-template/internal/helpers"
	"sort"
	"strings"
)

type Counts struct {
	letter byte
	count  int64
}

func PartOne(lines []string) string {
	result := int64(0)
	for _, line := range lines {
		left, right, found := strings.Cut(line, "[")
		if !found {
			continue
		}
		right = right[:len(right)-1]
		parts := strings.Split(left, "-")
		id := h.ToInt(parts[len(parts)-1])
		parts = parts[:len(parts)-1]
		counts := make([]Counts, int('z'-'a'+1))
		for i := range len(counts) {
			counts[i].letter = byte(i + 'a')
		}
		for _, part := range parts {
			for _, ch := range part {
				counts[ch-'a'].count++
			}
		}
		sort.Slice(counts, func(i, j int) bool {
			return counts[i].count > counts[j].count || counts[i].count == counts[j].count && counts[i].letter < counts[j].letter
		})
        valid := true
        for i := range 5 {
            if counts[i].letter != right[i] {
                valid = false
                break
            }
        }
        if valid {
            result += id
        }

	}
	return h.ToString(result)
}

func PartTwo(lines []string) string {
	for _, line := range lines {
		left, right, found := strings.Cut(line, "[")
		if !found {
			continue
		}
		right = right[:len(right)-1]
		parts := strings.Split(left, "-")
		id := h.ToInt(parts[len(parts)-1])
		parts = parts[:len(parts)-1]
		counts := make([]Counts, int('z'-'a'+1))
		for i := range len(counts) {
			counts[i].letter = byte(i + 'a')
		}
		for _, part := range parts {
			for _, ch := range part {
				counts[ch-'a'].count++
			}
		}
		sort.Slice(counts, func(i, j int) bool {
			return counts[i].count > counts[j].count || counts[i].count == counts[j].count && counts[i].letter < counts[j].letter
		})
        valid := true
        for i := range 5 {
            if counts[i].letter != right[i] {
                valid = false
                break
            }
        }
        if !valid {
			continue
        }
		result := ""
		for _, ch := range left {
			if ch >= 'a' && ch <= 'z'  {
				result += string('a' + (int(ch - 'a') +int(id)) % int('z' - 'a' + 1))
			} else {
				result += string(ch)
			}
		}
		if strings.HasPrefix(result, "northpole") {
			return h.ToString(id)
		}
	}
	return h.ToString(0)
}
