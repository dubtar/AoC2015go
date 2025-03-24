package day17

import (
	"crypto/md5"
	"encoding/hex"
	h "go-aoc-template/internal/helpers"

	"github.com/idsulik/go-collections/v2/deque"
)

type Data struct {
	position h.Coords
	path string
}

func PartOne(lines []string) string {
	key := lines[0]
	Min := int64(1)
	Max := int64(4)
	Target := h.Coords{X: Max, Y: Max}
	Dirs := []rune{'U', 'D', 'L', 'R'}
	queue := deque.New[Data](-1)
	queue.PushBack(Data{position: h.Coords{X: Min, Y: Min}, path: ""})
	for !queue.IsEmpty() {
		cur, _ := queue.PopFront()
		if cur.position == Target {
			return cur.path
		}
		hash := md5.Sum([]byte(key + cur.path))
		hex := hex.EncodeToString(hash[:])
		for i, val := range hex[:4] {
			if val < 'b' {
				continue
			}
			dir := Dirs[i]
			nextPos := cur.position.Plus(h.DirectionsUDLR[dir])
			if nextPos.X < Min || nextPos.Y < Min || nextPos.X > Max || nextPos.Y > Max {
				continue
			}
			queue.PushBack(Data{nextPos, cur.path+string(dir)})
		}
	}
	return "Not found"
}

func PartTwo(lines []string) string {
	key := lines[0]
	Min := int64(1)
	Max := int64(4)
	Target := h.Coords{X: Max, Y: Max}
	Dirs := []rune{'U', 'D', 'L', 'R'}
	queue := deque.New[Data](-1)
	queue.PushBack(Data{position: h.Coords{X: Min, Y: Min}, path: ""})
	longest := Data{}
	for !queue.IsEmpty() {
		cur, _ := queue.PopFront()
		if cur.position == Target {
			longest = cur
			continue
		}
		hash := md5.Sum([]byte(key + cur.path))
		hex := hex.EncodeToString(hash[:])
		for i, val := range hex[:4] {
			if val < 'b' {
				continue
			}
			dir := Dirs[i]
			nextPos := cur.position.Plus(h.DirectionsUDLR[dir])
			if nextPos.X < Min || nextPos.Y < Min || nextPos.X > Max || nextPos.Y > Max {
				continue
			}
			queue.PushBack(Data{nextPos, cur.path+string(dir)})
		}
	}
	return h.ToString(len(longest.path))
}
