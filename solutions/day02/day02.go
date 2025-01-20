package day02

import (
	h "go-aoc-template/internal/helpers"
)

func PartOne(lines []string) string {
	const S = 2
	pos := h.Coords{X: 1, Y: 1}
	result := ""
	for _, line := range lines {
		for _, dir := range line {
			pos.Add(h.DirectionsUDLR[dir])
			pos.X = min(max(pos.X, 0), S)
			pos.Y = min(max(pos.Y, 0), S)
		}
		result += h.ToString(pos.X + pos.Y*3 + 1)
	}

	return result
}

func getNum(pos h.Coords) (rune, bool) {
	switch pos.Y {
	case 0:
		if pos.X == 2 {
			return '1', true
		}
	case 1:
		if pos.X > 0 && pos.X < 4 {
			return rune(int64('1') + pos.X), true
		}
	case 2:
		if pos.X >= 0 && pos.X < 5 {
			return rune(int64('5') + pos.X), true
		}
	case 3:
		if pos.X > 0 && pos.X < 4 {
			return rune(int64('A') + pos.X - 1), true
		}
	case 4:
		if pos.X == 2 {
			return 'D', true
		}
	}
	return ' ', false
}
func PartTwo(lines []string) string {
	const S = 5
	field := make([][]rune, S)
	for i := range S {
		field = append(field, make([]rune, S))
		field[i] = make([]rune, S)
	}

	pos := h.Coords{X: 0, Y: 2}
	result := ""
	for _, line := range lines {
		for _, dir := range line {
			newpos := pos
			newpos.Add(h.DirectionsUDLR[dir])
			if _, ok := getNum(newpos); ok {
				pos = newpos
			}
		}
		res, _ := getNum(pos)
		result += string(res)
	}

	return result
}
