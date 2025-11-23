package day24

import (
	"fmt"
	h "go-aoc-template/internal/helpers"
	"math"

	"github.com/idsulik/go-collections/v2/deque"
)

func parse(lines []string) (result [][]rune, targets []h.Coords) {
	result = make([][]rune, 0, len(lines))
	targets = make([]h.Coords, 9)
	max_targets := 0
	for i, input_line := range lines {
		line := make([]rune, len(input_line))
		for j, c := range input_line {
			line[j] = c
			if c >= '0' && c <= '9' {
				ind := h.ToInt(string(c))
				targets[ind] = h.Coords{X: int64(i), Y: int64(j)}
				max_targets = max(max_targets, int(ind))
			}
		}
		result = append(result, line)
	}
	targets = targets[:max_targets+1]
	return
}

func FindPath(roofMap [][]rune, start h.Coords, target h.Coords) int {
	type Item struct {
		pos    h.Coords
		result int
	}
	queue := deque.New[Item](0)

	visited := map[h.Coords]bool{}
	queue.PushBack(Item{pos: start, result: 0})
	for queue.Len() > 0 {
		item, ok := queue.PopFront()
		if !ok {
			panic("queue is empty")
		}
		for _, dir := range h.DirectionsHV {
			newPos := item.pos.Plus(dir)
			if newPos == target {
				return item.result + 1
			}
			if newPos.X < 0 || newPos.Y < 0 || newPos.X >= int64(len(roofMap)) || newPos.Y >= int64(len(roofMap[0])) {
				continue
			}
			if roofMap[newPos.X][newPos.Y] == '#' || visited[newPos] {
				continue
			}
			visited[newPos] = true
			queue.PushBack(Item{pos: newPos, result: item.result + 1})
		}
	}
	panic(fmt.Sprintf("path not found, start: %v, target: %v", start, target))
}

func buildPaths(roofMap [][]rune, targets []h.Coords) (paths [][]int) {
	paths = make([][]int, len(targets))
	for sourceIndex, source := range targets {
		paths[sourceIndex] = make([]int, len(targets))
		for targetIndex, target := range targets {
			if sourceIndex == targetIndex {
				paths[sourceIndex][targetIndex] = -1
				continue
			}
			paths[sourceIndex][targetIndex] = FindPath(roofMap, source, target)
		}
	}
	return paths
}

func PartOne(lines []string) string {
	roofMap, targets := parse(lines)
	paths := buildPaths(roofMap, targets)
	// полный перебор
	best := math.MaxInt64
	indexes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}[:len(targets)-1]
	for _, perm := range h.Permutations(indexes) {
		path, prev := 0, 0
		for _, i := range perm {
			path += paths[prev][i]
			prev = i
		}
		best = min(best, path)
	}

	return h.ToString(best)
}
func PartTwo(lines []string) string {
	roofMap, targets := parse(lines)
	paths := buildPaths(roofMap, targets)
	// полный перебор
	best := math.MaxInt64
	indexes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}[:len(targets)-1]
	for _, perm := range h.Permutations(indexes) {
		path, prev := 0, 0
		for _, i := range perm {
			path += paths[prev][i]
			prev = i
		}
		path += paths[prev][0]
		best = min(best, path)
	}

	return h.ToString(best)
}
