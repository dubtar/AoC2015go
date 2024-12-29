package day17

import (
	h "go-aoc-template/internal/helpers"
	"strconv"
)

func find(target int64, conts []int64, path int, paths map[int]int) int {
	if len(conts) == 0 || target < 0 {
		return 0
	}
	res1 := find(target, conts[1:], path, paths)
	res2 := 0
	if conts[0] == target {
		res2 = 1
		paths[path+1] += 1
	} else {
		res2 = find(target-conts[0], conts[1:], path+1, paths)
	}
	return res1 + res2
}
func PartOneTargeted(lines []string, target int64) (int, int) {
	cont := make([]int64, len(lines))
	for i, line := range lines {
		cont[i] = h.ToInt(line)
	}
	paths := make(map[int]int)
	res := find(target, cont, 0, paths)
	minKey := len(cont)
	for k := range paths {
		minKey = min(minKey, k)
	}
	return res, paths[minKey]
}
func PartOne(lines []string) string {
	res, _ := PartOneTargeted(lines, 150)
	return strconv.Itoa(res)
}

func PartTwo(lines []string) string {
	_, res := PartOneTargeted(lines, 150)
	return strconv.Itoa(res)
}
