package day11

import (
	"github.com/idsulik/go-collections/v2/deque"
	h "go-aoc-template/internal/helpers"
	"slices"
	"strings"
)

type Floors [][]string

func (d *Data) Hash() string {
	hash := h.ToString(d.elevator)
	for _, v := range d.floors {
		slices.Sort(v)
		hash += strings.Join(v, ",") + ";"
	}
	return hash
}

type Data struct {
	floors   Floors
	elevator int
	steps    int
	log      []string
}

func (f Floors) Clone() Floors {
	res := make([][]string, 0, len(f))
	for _, v := range f {
		res = append(res, slices.Clone(v))
	}
	return res
}

func (f Floors) Valid() bool {
	for _, v := range f {
		for _, e := range v {
			t := e[0]
			if e[1] == 'M' {
				hasOwn := false
				hasOther := false
				for _, ee := range v {
					if ee[1] == 'G' {
						if ee[0] == t {
							hasOwn = true
							break
						} else {
							hasOther = true
						}
					}
				}
				if !hasOwn && hasOther {
					return false
				}
			}
		}
	}
	return true
}

var Dirs = []int{-1, 1}

func PartOne(lines []string) string {
	var TotalItems = 10
	const Floors = 4
	var start = [][]string{{"SG", "SM", "PG", "PM"}, {"TG", "RG", "RM", "CG", "CM"}, {"TM"}, {}}
	if len(lines) == 1 {
		TotalItems = 4
		start = [][]string{{"HM", "LM"}, {"HG"}, {"LG"}, {}}
	}
	// queue := list.New()
	var queue = deque.New[*Data](100)
	var visited = make(map[string]bool)
	queue.PushBack(&Data{floors: start, steps: 0, elevator: 0})
	for {
		cur, ok := queue.PopFront()
		if !ok {
			return "Not found"
		}
		// if queue.Len() == 0 {
		// 	return "Not found"
		// }
		// cur := queue.Front().Value.(*Data)
		// queue.Remove(queue.Front())
		if len(cur.floors[len(start)-1]) == TotalItems {
			return h.ToString(cur.steps)
		}
		for _, dir := range Dirs {
			nextElevator := cur.elevator + dir
			if nextElevator < 0 || nextElevator >= len(start) {
				continue
			}
			for i, v := range cur.floors[cur.elevator] {
				newFloors := cur.floors.Clone()
				newFloors[cur.elevator] = append(newFloors[cur.elevator][:i], newFloors[cur.elevator][i+1:]...)
				newFloors[nextElevator] = append(newFloors[nextElevator], v)
				if newFloors.Valid() {
					d := Data{floors: newFloors, steps: cur.steps + 1, elevator: nextElevator}
					hash := d.Hash()
					if !visited[hash] {
						visited[hash] = true
						queue.PushBack(&d)
					}
				}
				for j, w := range newFloors[cur.elevator] {
					newNewFloors := newFloors.Clone()
					newNewFloors[cur.elevator] = append(newNewFloors[cur.elevator][:j], newNewFloors[cur.elevator][j+1:]...)
					newNewFloors[nextElevator] = append(newNewFloors[nextElevator], w)
					if newNewFloors.Valid() {
						vv := Data{floors: newNewFloors, steps: cur.steps + 1, elevator: nextElevator}
						hash := vv.Hash()
						if !visited[hash] {
							visited[hash] = true
							queue.PushBack(&vv)
						}
					}
				}
			}
		}
	}
}

func PartTwo(lines []string) string {
	var start = [][]string{{"DM", "DG", "EM", "EG", "SG", "SM", "PG", "PM"}, {"TG", "RG", "RM", "CG", "CM"}, {"TM"}, {}}
	var totalItems = 0
	for _, v := range start {
		totalItems += len(v)
	}
	var queue = deque.New[*Data](100)
	var visited = make(map[string]bool)
	queue.PushBack(&Data{floors: start, steps: 0, elevator: 0})
	for {
		cur, ok := queue.PopFront()
		if !ok {
			return "Not found"
		}
		if len(cur.floors[len(start)-1]) == totalItems {
			return h.ToString(cur.steps)
		}
		for _, dir := range Dirs {
			nextElevator := cur.elevator + dir
			if nextElevator < 0 || nextElevator >= len(start) {
				continue
			}
			for i, v := range cur.floors[cur.elevator] {
				newFloors := cur.floors.Clone()
				newFloors[cur.elevator] = append(newFloors[cur.elevator][:i], newFloors[cur.elevator][i+1:]...)
				newFloors[nextElevator] = append(newFloors[nextElevator], v)
				if newFloors.Valid() {
					d := Data{floors: newFloors, steps: cur.steps + 1, elevator: nextElevator}
					hash := d.Hash()
					if !visited[hash] {
						visited[hash] = true
						queue.PushBack(&d)
					}
				}
				for j, w := range newFloors[cur.elevator] {
					newNewFloors := newFloors.Clone()
					newNewFloors[cur.elevator] = append(newNewFloors[cur.elevator][:j], newNewFloors[cur.elevator][j+1:]...)
					newNewFloors[nextElevator] = append(newNewFloors[nextElevator], w)
					if newNewFloors.Valid() {
						vv := Data{floors: newNewFloors, steps: cur.steps + 1, elevator: nextElevator}
						hash := vv.Hash()
						if !visited[hash] {
							visited[hash] = true
							queue.PushBack(&vv)
						}
					}
				}
			}
		}
	}
}
