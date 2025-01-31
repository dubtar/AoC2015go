package day05

import (
	"crypto/md5"
	"encoding/hex"
	h "go-aoc-template/internal/helpers"
)

func PartOne(lines []string) string {
	i := int64(0)
	result := []byte{}
	for {
		hash := md5.Sum([]byte(lines[0] + h.ToString(i)))
		hex := hex.EncodeToString(hash[:])
		if hex[:5] == "00000" {
			result = append(result, hex[5])
            if len(result) == 8 {
                break
            }
		}
		i++
	}
	return string(result)
}

func PartTwo(lines []string) string {
	i := int64(0)
	result := make([]byte, 8)
	for {
		hash := md5.Sum([]byte(lines[0] + h.ToString(i)))
		hex := hex.EncodeToString(hash[:])
		if hex[:5] == "00000" {
			pos := hex[5] - '0'
			if pos < 8 && result[pos] == 0 {
				result[pos] = hex[6]
			}
			done := true
			for _, a := range result {
				if a == 0 {
					done = false
					break
				}
			}
			if done {
				break
			}
		}
		i++
	}
	return string(result)
}
