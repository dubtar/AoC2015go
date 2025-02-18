package day13

import (
	h "go-aoc-template/internal/helpers"

	"github.com/idsulik/go-collections/v2/deque"
)

func IsWall(c h.Coords, num int64) bool {
	x, y := c.X, c.Y
	t := x*x + 3*x + 2*x*y + y + y*y + num
	return h.BitCount(t)%2 == 1
}

type Data struct {
	h.Coords
	steps int
}

func PartOne(lines []string) string {
	num := h.ToInt(lines[0])
	target := h.Coords{X: 31, Y: 39}
	if lines[0] == "10" {
		target = h.Coords{X: 7, Y: 4}
	}
	queue := deque.New[Data](-1)
	queue.PushBack(Data{Coords: h.Coords{X: 1, Y: 1}, steps: 0})
	visited := make(map[h.Coords]bool)
	visited[h.Coords{X: 1, Y: 1}] = true
	for {
		cur, ok := queue.PopFront()
		if !ok {
			return "Not found"
		}
		if cur.Coords == target {
            print("\n")
            Display(visited, num)
			return h.ToString(cur.steps)
		}
		for _, dir := range h.DirectionsUDLR {
			next := h.Coords{X: cur.X + dir.X, Y: cur.Y + dir.Y}
			if next.X < 0 || next.Y < 0 {
				continue
			}
			if !visited[next] && !IsWall(next, num) {
				visited[next] = true
				queue.PushBack(Data{Coords: next, steps: cur.steps + 1})
			}
		}

	}
}
func Display(visited map[h.Coords]bool, num int64) {
    for y := int64(0); y < 50; y++ {
        for x := int64(0); x < 50; x++ {
            if visited[h.Coords{X: x, Y: y}] {
                print("O")
            } else if IsWall(h.Coords{X: x, Y: y}, num) {
                print("#")
            } else {
                print(".")
            }
        }
        print("\n")
    }
}

func PartTwo(lines []string) string {
	num := h.ToInt(lines[0])

	queue := deque.New[Data](-1)
	queue.PushBack(Data{Coords: h.Coords{X: 1, Y: 1}, steps: 0})
	visited := make(map[h.Coords]bool)
	visited[h.Coords{X: 1, Y: 1}] = true
	for {
		cur, ok := queue.PopFront()
		if !ok {
			return "Not found"
		}
		if cur.steps >= 50 {
            Display(visited, num)
			return h.ToString(len(visited))
		}
		for _, dir := range h.DirectionsUDLR {
			next := h.Coords{X: cur.X + dir.X, Y: cur.Y + dir.Y}
			if next.X < 0 || next.Y < 0 {
				continue
			}
			if !visited[next] && !IsWall(next, num) {
				visited[next] = true
				queue.PushBack(Data{Coords: next, steps: cur.steps + 1})
			}
		}
	}
}
