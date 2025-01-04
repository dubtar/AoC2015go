package day24

import (
	"fmt"
	h "go-aoc-template/internal/helpers"
	"math"
	"slices"

	"github.com/idsulik/go-collections/v2/deque"
)

func parse(lines []string) ([]int64, int64) {
	result := make([]int64, len(lines))
	sum := int64(0)
	for i, line := range lines {
		result[i] = h.ToInt(line)
		sum += result[i]
	}
	return result, sum
}

func PartOne(lines []string) string {
	return Find(lines, 3)
}
func Find(lines []string, parts int) string {
	boxes, sum := parse(lines)
	slices.Sort(boxes)
	slices.Reverse(boxes)
	target := sum / int64(parts)
	if target*int64(parts) != sum {
		panic(fmt.Sprint(sum, "not / 3"))
	}
	bestLen := len(boxes)
	bestQ := int64(math.MaxInt64)
	queue := deque.New[[]int](10)
	queue.PushBack([]int{0})
	for {
		path, ok := queue.PopFront()
		if !ok {
			break
		}
		if len(path) >= bestLen {
			continue
		}
		curSum := int64(0)
		lastind := path[len(path)-1]
		for _, ind := range path {
			curSum += boxes[ind]
		}
		for i := lastind + 1; i < len(boxes); i++ {
			nextSum := curSum + boxes[i]
			if nextSum > target {
				continue
			}
			if nextSum == target {
				bestLen = min(bestLen, len(path)+1)
				q := boxes[i]
				for _, ind := range path {
					q *= boxes[ind]
				}
				bestQ = min(bestQ, q)
			}
			newpath := make([]int, len(path)+1)
			copy(newpath, path)
			newpath[len(newpath)-1] = i
			queue.PushBack(newpath)
		}
	}
	return h.ToString(bestQ)
}

func PartTwo(lines []string) string {
	return Find(lines, 4)
}
