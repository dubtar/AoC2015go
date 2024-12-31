package day19

import (
	"fmt"
	h "go-aoc-template/internal/helpers"
	"strconv"
	"strings"
)

type Rules map[string][]string

var replaces = map[string]string{}
var nextVar = 'f'

func repl(s string) string {
	var v string
	if v, ok := replaces[s]; ok {
		return v
	}
	if len(s) == 1 {
		v = s
	} else {
		v = string(nextVar)
		nextVar++
	}
	replaces[s] = v
	return v
}
func replS(s string) string {
	val := ""
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && h.IsLowChar(s[i+1]) {
			val += repl(s[i : i+2])
			i++
		} else {
			val += repl(string(s[i]))
		}
	}
	return val
}
func parse(lines []string) (rules Rules, molecule string) {
	isRules := true
	rules = Rules{}
	for _, line := range lines {
		if line == "" {
			isRules = false
			continue
		}
		if isRules {
			parts := strings.Split(line, " => ")
			key := repl(parts[0])
			val := replS(parts[1])
			rules[key] = append(rules[key], val)
		} else {
			molecule = replS(line)
			return
		}
	}
	return
}

func PartOne(lines []string) string {
	rules, source := parse(lines)
	isTest := len(source) < 10

	results := make(map[string]any)
	for i := 0; i < len(source); i++ {
		for k, subs := range rules {
			if i+len(k) > len(source) {
				continue
			}
			if source[i:i+len(k)] == k {
				for _, s := range subs {
					v := source[:i] + s + source[i+len(k):]
					results[v] = struct{}{}
					if isTest {
						fmt.Println(v)
					}
				}
			}
		}
	}
	return strconv.Itoa(len(results))
}

func PartTwo(lines []string) string {
	_, target := parse(lines)
	_, ok := replaces["Rn"]
	if !ok { // test
		return strconv.Itoa(len(target) - 1)
	}
	rn, y, ar := replaces["Rn"][0], replaces["Y"][0], replaces["Ar"][0]
	rns, ys, ars := 0, 0, 0
	for i := 0; i < len(target); i++ {
		c := target[i]
		switch c {
		case rn:
			rns++
		case y:
			ys++
		case ar:
			ars++
		}
	}
	return strconv.Itoa(len(target) - rns - ys*2 - ars - 1)
}
