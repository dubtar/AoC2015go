package day20

import (
	// h "go-aoc-template/internal/helpers"
	"slices"
	"strconv"
	"strings"
)

type pair struct {
	start int
	end   int
}

func PartOne(lines []string) string {
	bans := make([]pair, 0, len(lines))
	for _, line := range lines {
		p := pair{}
		ind := strings.Index(line, "-")
		p.start, _ = strconv.Atoi(line[0:ind])
		p.end, _ = strconv.Atoi(line[ind+1:])
		bans = append(bans, p)
	}
	slices.SortFunc(bans, func(a, b pair) int {
		if a.start != b.start {
			return a.start - b.start
		}
		return a.end - b.end
	})
	lastBan := bans[0].end
	for _, ban := range bans[1:] {
		if ban.start > lastBan+1 {
			return strconv.Itoa(lastBan + 1)
		}
		lastBan = max(lastBan, ban.end)
	}
	return "failed"
}

func PartTwo(lines []string) string {
	bans := make([]pair, 0, len(lines))
	maxValue := 1 << 32
	if len(lines) < 10 {
		maxValue = 10
	}
	for _, line := range lines {
		p := pair{}
		ind := strings.Index(line, "-")
		p.start, _ = strconv.Atoi(line[0:ind])
		p.end, _ = strconv.Atoi(line[ind+1:])
		bans = append(bans, p)
	}
	slices.SortFunc(bans, func(a, b pair) int {
		if a.start != b.start {
			return a.start - b.start
		}
		return a.end - b.end
	})
	result := 0
	lastBan := bans[0].end
	for _, ban := range bans[1:] {
		if ban.start > lastBan+1 {
			result += ban.start - lastBan - 1
		}
		lastBan = max(lastBan, ban.end)
	}
	if lastBan < maxValue - 1 {
		result += maxValue - lastBan - 1
	}
	return strconv.Itoa(result)
}
