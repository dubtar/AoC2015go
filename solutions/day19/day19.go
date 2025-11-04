package day19

import (
	"fmt"
	h "go-aoc-template/internal/helpers"
	"math"
	"strconv"
)

func PartOne(lines []string) string {
	num, _ := strconv.Atoi(lines[0])
	arr := make([]int, num)
	for i := range num {
		arr[i] = 1
	}
	cur := 0
	i := cur
	isStealing := true
	for num > 1 {
		i++
		if i == len(arr) {
			i = 0
		}
		if arr[i] == 0 {
			continue
		}
		if isStealing {
			// забираем
			arr[cur] += arr[i]
			arr[i] = 0
			num--
			isStealing = false
		} else {
			cur = i
			isStealing = true
		}
	}

	return strconv.Itoa(cur + 1)
}

func PartTwo2(lines []string) string {
	// моё решение
	num, _ := strconv.Atoi(lines[0])
	t := 10000
	if num < 100 {
		t = 2
	}
	arr := h.NewChainedArray[int](t)
	for i := range num {
		arr.Append(i + 1)
	}
	cur := 0
	for num > 1 {
		target := (cur + num/2) % num
		arr.Delete(target)
		num--
		fmt.Printf("\r n=%v", num)
		if target > cur {
			cur++
		}
		cur = cur % num
	}

	return strconv.Itoa(arr.Get(0))
}

func PartTwo(lines []string) string {
	// подсмотренное решение
	num, _ := strconv.Atoi(lines[0])
	arr := make([]int, num)
	for i := range num {
		arr[i] = i + 1
	}

	for len(arr) > 1 {
		num := len(arr)
		eliminated := 0
		limit := int(math.Ceil(float64(num) / 3))
		for i := range limit {
			across := i + eliminated + (num / 2)
			arr[across] = 0
			num -= 1
			eliminated += 1
		}
		result := make([]int, 0, num)
		for _, c := range arr[limit:] {
			if c != 0 {
				result = append(result, c)
			}
		}
		for _, c := range arr[:limit] {
			if c != 0 {
				result = append(result, c)
			}
		}
		arr = result
	}

	return strconv.Itoa(arr[0])
}
