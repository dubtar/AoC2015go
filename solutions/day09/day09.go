package day09

import (
	h "go-aoc-template/internal/helpers"
	"strings"
)

func PartOne(lines []string) string {
    result := int64(0)
    line := lines[0]
    for i := 0; i < len(line); i++ {
        if line[i] == '(' {
            indx := strings.Index(line[i+1:], "x") + i + 1
            indEnd := strings.Index(line[indx+1:], ")") + indx + 1
            count := h.ToInt(line[i+1:indx])
            repeat := h.ToInt(line[indx+1:indEnd])
            result += count*repeat
            i = indEnd + int(count)
        } else {
            result += 1
        }
    }
    return h.ToString(result)
}

func decompress(line string) int64 {
    result := int64(0)
    for i := 0; i < len(line); i++ {
        if line[i] == '(' {
            indx := strings.Index(line[i+1:], "x") + i + 1
            indEnd := strings.Index(line[indx+1:], ")") + indx + 1
            count := h.ToInt(line[i+1:indx])
            repeat := h.ToInt(line[indx+1:indEnd])
            result += decompress(line[indEnd+1:indEnd+1+int(count)])*repeat
            i = indEnd + int(count)
        } else {
            result += 1
        }
    }
    return result
}

func PartTwo(lines []string) string {
    return h.ToString(decompress(lines[0]))
}
