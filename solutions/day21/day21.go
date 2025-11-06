package day21

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	// h "go-aoc-template/internal/helpers"
)

func PartOne(lines []string) string {
	value := strings.Split("abcdefgh", "")
	isDebug := false
	if len(lines) < 10 {
		isDebug = true
		value = strings.Split("abcde", "")
	}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) == 0 {
			continue
		}
		switch parts[0] {
		case "swap":
			indLeft, indRight := 0, 0
			switch parts[1] {
			case "position":
				indLeft, _ = strconv.Atoi(parts[2])
				indRight, _ = strconv.Atoi(parts[5])
			case "letter":
				indLeft = slices.Index(value, parts[2])
				indRight = slices.Index(value, parts[5])
			}

			value[indLeft], value[indRight] = value[indRight], value[indLeft]
		case "reverse":
			indLeft, _ := strconv.Atoi(parts[2])
			indRight, _ := strconv.Atoi(parts[4])
			for indLeft < indRight {
				value[indLeft], value[indRight] = value[indRight], value[indLeft]
				indLeft++
				indRight--
			}
		case "rotate":
			ind := 0
			switch parts[1] {
			case "left":
				ind, _ = strconv.Atoi(parts[2])
			case "right":
				ind, _ = strconv.Atoi(parts[2])
				ind = len(value) - ind
			case "based":
				ind = slices.Index(value, parts[6])
				if ind >= 4 {
					ind += 1
				}
				ind++
				ind = len(value) - (ind % len(value))
			}
			value = append(value[ind:], value[:ind]...)
		case "move":
			indLeft, _ := strconv.Atoi(parts[2])
			indRight, _ := strconv.Atoi(parts[5])
			dir := 1
			if indLeft > indRight {
				dir = -1
			}
			l := value[indLeft]
			for ; indLeft != indRight; indLeft += dir {
				value[indLeft] = value[indLeft+dir]
			}
			value[indRight] = l
		}
		if isDebug {
			fmt.Printf("%s: %s\n", line, strings.Join(value, ""))
		}
	}
	return strings.Join(value, "")
}

func PartTwo(lines []string) string {
	value := strings.Split("fbgdceah", "")
	// 01234567 : 0 -> 1 ( 0 + 0 1)
	// b in aBcdefg : 1 -> 3 (1 + 1 + 1)
	// c in 01c34567 : 2 -> 5 ( 2 + 2 + 1)
	// d in 012d4567 : 3 -> 7 (7=3 + 3 + 1)
	// e in 0123e567 : 4 -> 2 (10=4+4 + 2)
	// f in 01234f67 : 5 -> 4 (12=5 + 5 + 2)
	// g in 012345g7 : 6 -> 6 (14=6 + 6 + 2)
	// h in 0123456h : 7 -> 0 (16=7 + 7 + 2)
	//
	// 01234 : 0 -> 1
	// 1 -> 3
	// 2 -> 0 (5= 2 + 2 + 1)
	// 3 -> 2
	// 4 -> 0 (10=4 + 4 + 2)
	rotates := []int{7, 0, 4, 1, 5, 2, 6, 3}
	if len(lines) < 10 {
		value = strings.Split("decab", "")
		rotates = []int{4, 0, 3, 1, -1}
	}
	for _, line := range slices.Backward(lines) {
		parts := strings.Split(line, " ")
		if len(parts) == 0 {
			continue
		}
		switch parts[0] {
		case "swap":
			indLeft, indRight := 0, 0
			switch parts[1] {
			case "position":
				indLeft, _ = strconv.Atoi(parts[2])
				indRight, _ = strconv.Atoi(parts[5])
			case "letter":
				indLeft = slices.Index(value, parts[2])
				indRight = slices.Index(value, parts[5])
			}

			value[indLeft], value[indRight] = value[indRight], value[indLeft]
		case "reverse":
			indLeft, _ := strconv.Atoi(parts[2])
			indRight, _ := strconv.Atoi(parts[4])
			for indLeft < indRight {
				value[indLeft], value[indRight] = value[indRight], value[indLeft]
				indLeft++
				indRight--
			}
		case "rotate":
			ind := 0
			switch parts[1] {
			case "left":
				ind, _ = strconv.Atoi(parts[2])
				ind = len(value) - ind // reverse
			case "right":
				ind, _ = strconv.Atoi(parts[2])
				// reverse
			case "based":
					ind = slices.Index(value, parts[6])
					ind = (len(value) + ind - rotates[ind]) % len(value) // reverse
			}
			value = append(value[ind:], value[:ind]...)
		case "move":
			indLeft, _ := strconv.Atoi(parts[2])
			indRight, _ := strconv.Atoi(parts[5])
			dir := -1 // reverse
			if indLeft > indRight {
				dir = 1
			}
			l := value[indRight]
			for ; indLeft != indRight; indRight += dir {
				value[indRight] = value[indRight+dir]
			}
			value[indRight] = l
		}
		fmt.Printf("%s: %s\n", line, strings.Join(value, ""))
	}
	return strings.Join(value, "")
}
