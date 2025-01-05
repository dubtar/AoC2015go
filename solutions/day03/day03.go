package day03

import (
	h "go-aoc-template/internal/helpers"
)

func PartOne(lines []string) string {
	pos := h.Coords{X: 0, Y: 0}
	visited := map[h.Coords]any{pos: struct{}{}}

	for _, dir := range lines[0] {
		pos.Add(h.DirectionsHV[dir])
		visited[pos] = struct{}{}
	}
	return h.ToString(len(visited))
}

func PartTwo(lines []string) string {
	pos1 := h.Coords{X: 0, Y: 0}
	pos2 := pos1
	visited := map[h.Coords]any{pos1: struct{}{}}

	turn2 := false
	for _, dir := range lines[0] {
		if turn2 {
			pos2.Add(h.DirectionsHV[dir])
			visited[pos2] = struct{}{}
		} else {
			pos1.Add(h.DirectionsHV[dir])
			visited[pos1] = struct{}{}
		}
		turn2 = !turn2
	}
	return h.ToString(len(visited))
}
