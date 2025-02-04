package day08

import (
	"fmt"
	h "go-aoc-template/internal/helpers"
	"strings"
)

var Width = int64(50)
var Height = int64(6)

func Print( screen [][]bool) {
    for _, row := range screen {
        for _, px := range row {
            if px {
                fmt.Print("#")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}

func PartOne(lines []string) string {
	screen := make([][]bool, Height)
	for i := int64(0); i < Height; i++ {
		screen[i] = make([]bool, Width)
	}
	for _, line := range lines {
		if strings.HasPrefix(line, "rect ") {
			val := strings.Split(line[len("rect "):], "x")
			for x := range h.ToInt(val[0]) {
				for y := range h.ToInt(val[1]) {
					screen[y][x] = true
				}
			}
		} else if strings.HasPrefix(line, "rotate row") {
			vals := strings.Split(line[len("rotate row y="):], " ")
			y := h.ToInt(vals[0])
			shift := h.ToInt(vals[2])
			screen[y] = append(screen[y][Width-shift:], screen[y][:Width-shift]...)
		} else if strings.HasPrefix(line, "rotate column") {
			vals := strings.Split(line[len("rotate column x="):], " ")
			x := h.ToInt(vals[0])
			shift := h.ToInt(vals[2])
            newcol := make([]bool, Height)
            for i := range Height {
                newcol[i] = screen[(Height + i - shift) % Height][x]
            }
            for i:= range Height {
                screen[i][x] = newcol[i]
            }
		}
	}
    result := int64(0)
    for y := range Height {
        for x := range Width {
            if  screen[y][x] {
                result++
            }
        }
    }
    Print(screen)

	return h.ToString(result)
}

func PartTwo(lines []string) string {
	return "RURUCEOEIL" // прочитать вывод
}
