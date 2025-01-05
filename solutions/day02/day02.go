package day02

import (
	"fmt"
	h "go-aoc-template/internal/helpers"
	"strings"
)

func PartOne(lines []string) string {
	area := int64(0)
	for _, line := range lines {
		parts := strings.Split(line, "x")
		l, w, h := h.ToInt(parts[0]), h.ToInt(parts[1]), h.ToInt(parts[2])
		area += 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
	}

	return fmt.Sprint(area)
}

func PartTwo(lines []string) string {
	ribbon := int64(0)
	for _, line := range lines {
		parts := strings.Split(line, "x")
		l, w, h := h.ToInt(parts[0]), h.ToInt(parts[1]), h.ToInt(parts[2])
		ribbon += 2*min(l+w, w+h, h+l) + l*w*h
	}

	return fmt.Sprint(ribbon)
}
