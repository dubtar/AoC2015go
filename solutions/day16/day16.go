package day16

import (
	h "go-aoc-template/internal/helpers"
	"strings"
)

func parse(lines []string) map[string]map[string]int64 {
	result := make(map[string]map[string]int64)
	for _, line := range lines {
		parts := strings.SplitN(line, ": ", 2)
		name := strings.Split(parts[0], " ")[1]
		result[name] = make(map[string]int64)
		for _, part := range strings.Split(parts[1], ", ") {
			parts := strings.Split(part, ": ")
			result[name][parts[0]] = h.ToInt(parts[1])
		}
	}
	return result
}

var target = map[string]int64{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func PartOne(lines []string) string {
	aunts := parse(lines)
	for name, aunt := range aunts {
		found := true
		for key, value := range aunt {
			if target[key] != value {
				found = false
				break
			}
		}
		if found {
			return name
		}
	}

	return "Not found"
}

func PartTwo(lines []string) string {
	aunts := parse(lines)

	for name, aunt := range aunts {
		found := true
		for key, value := range aunt {
			switch key {
			case "cats", "trees":
				if target[key] >= value {
					found = false
					break
				}
			case "pomeranians", "goldfish":
				if target[key] <= value {
					found = false
					break
				}
			default:
				if target[key] != value {
					found = false
					break
				}
			}
		}
		if found {
			return name
		}
	}

	return "Not found"
}
