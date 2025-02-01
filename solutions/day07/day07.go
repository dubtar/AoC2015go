package day07

import (
	h "go-aoc-template/internal/helpers"
)

func supportsTLS(s string) bool {
    insideBrackets := false
    result := false
    for i := 0; i < len(s) - 3; i++ {
        c := s[i]
        if c == '[' {
            insideBrackets = true
            continue
        } else if c == ']' {
            insideBrackets = false
            continue
        }
        if byte(c) != s[i+1] && byte(c) == s[i+3] && s[i+1] == s[i+2] {
            if insideBrackets {
                return false
            }
            result = true
        }
    }
    return result
}


func PartOne(lines []string) string {
	result := 0
	for _, line := range lines {
		if supportsTLS(line) {
			result++
		}
	}
	return h.ToString(result)
}

func supportsSSL(s string) bool {
    insides := map[string]bool{}
    outsides := map[string]bool{}
    insideBrackets := false
    for i:= 0; i < len(s) - 2; i++ {
        if s[i+2] == '[' {
            insideBrackets = true
            i+=2
            continue
        } else if s[i+2] == ']' {
            insideBrackets = false
            i+=2
            continue
        }
        if s[i] == s[i+2] && s[i] != s[i+1] {
            if insideBrackets {
                insides[s[i:i+3]] = true
            } else {
                outsides[s[i:i+3]] = true
            }
        }
    }
    for k := range outsides {
        if _, ok := insides[k[1:3]+k[1:2]]; ok {
            return true
        }
    }
    return false
}
func PartTwo(lines []string) string {
	result := 0
	for _, line := range lines {
		if supportsSSL(line) {
			result++
		}
	}
	return h.ToString(result)
}
