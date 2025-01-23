package day03

import (
	h "go-aoc-template/internal/helpers"
	"sort"
	"strings"
)

func PartOne(lines []string) string {
	result := int64(0)
	for _, line := range lines {
		// extract nums from line
		nums := h.ToInts(strings.Split(strings.TrimSpace(line), " "))
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		if nums[0]+nums[1] > nums[2] {
			result++
		}
	}

	return h.ToString(result)
}

func PartTwo(lines []string) string {
	result := int64(0)
	for i := 0; i < len(lines)-2; i += 3 {
		part := make([][]int64, 3)
		for j := range 3 {
			part[j] = h.ToInts(strings.Split(strings.TrimSpace(lines[i+j]), " "))
		}

		for j := range 3 {
			nums := []int64{part[0][j], part[1][j], part[2][j]}
			sort.Slice(nums, func(i2, j2 int) bool {
				return nums[i2] < nums[j2]
			})
			if nums[0]+nums[1] > nums[2] {
				result++
			}
		}
	}

	return h.ToString(result)
}
