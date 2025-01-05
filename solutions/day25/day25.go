package day25

import (
	"fmt"
	h "go-aoc-template/internal/helpers"
	"regexp"
)

func PartOne(lines []string) string {
	r := regexp.MustCompile("(\\d+)")
	nums := r.FindAll([]byte(lines[0]), 2)
	target_row := h.ToInt(string(nums[0]))
	target_col := h.ToInt(string(nums[1]))
	row, col := int64(1), int64(1)
	value := uint64(20151125)
	for row != target_row || col != target_col {
		row--
		col++
		if row == 0 {
			row = col
			col = 1
		}
		value = (value * 252533) % 33554393
	}
	return fmt.Sprint(value)
}

func PartTwo(lines []string) string {
	return "TODO"
}
