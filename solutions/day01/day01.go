package day01

import (
	"strings"

	h "go-aoc-template/internal/helpers"
)

func PartOne(lines []string) string {
	instructions := strings.Split(lines[0], ", ")
	direction := '^'
	pos := h.Coords{X: 0, Y: 0}
	for _, instruction := range instructions {
		switch instruction[0] {
		case 'L':
			// turn left
			direction = h.TurnLeft(direction)
		case 'R':
			// turn right
			direction = h.TurnRight(direction)
		}
		steps := h.ToInt(instruction[1:])
		pos.Add(h.DirectionsHV[direction].Mul(steps))
	}

	return h.ToString(h.Abs(pos.X) + h.Abs(int64(pos.Y)))
}

func PartTwo(lines []string) string {
	instructions := strings.Split(lines[0], ", ")
	direction := '^'
	pos := h.Coords{X: 0, Y: 0}
	visits := make(map[h.Coords]any)
	for _, instruction := range instructions {
		switch instruction[0] {
		case 'L':
			// turn left
			direction = h.TurnLeft(direction)
		case 'R':
			// turn right
			direction = h.TurnRight(direction)
		}
		for range h.ToInt(instruction[1:]) {
			pos.Add(h.DirectionsHV[direction])
			if _, visited := visits[pos]; visited {
				return h.ToString(h.Abs(pos.X) + h.Abs(int64(pos.Y)))
			}
			visits[pos] = struct{}{}
		}
	}

	return "Not found"
}
