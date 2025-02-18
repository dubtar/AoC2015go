package day14

import (
	"crypto/md5"
	"encoding/hex"
	h "go-aoc-template/internal/helpers"
)

func GetTriplet(s string) byte {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return s[i]
		}
	}
	return 0
}

func GetFivplets(hash string) []byte {
	result := []byte{}
	for i := 0; i < len(hash)-4; i++ {
		if hash[i] == hash[i+1] && hash[i] == hash[i+2] && hash[i] == hash[i+3] && hash[i] == hash[i+4] {
			result = append(result, hash[i])
			i += 4
		}
	}
	return result
}

type Key struct {
	index  int64
	symbol byte
}

func Do(pref string, extra_cycles int) string {
	const Target = 64
	found_keys := 0
	potential_keys := map[byte][]Key{}
	for index := int64(0); true; index++ {
		md := md5.Sum([]byte(pref + h.ToString(index)))
		hash := hex.EncodeToString(md[:])
		for i := 0; i < extra_cycles; i++ {
			md := md5.Sum([]byte(hash))
			hash = hex.EncodeToString(md[:])
		}
		potential_key := GetTriplet(hash)
		if potential_key != 0 {
			potential_keys[potential_key] = append(potential_keys[potential_key], Key{index, potential_key})
		}
		for _, fiveplet := range GetFivplets(hash) {
			for _, key := range potential_keys[fiveplet] {
				if key.index+1000 >= index && key.index < index {
					found_keys++
					print(found_keys, ": Found key ", string(rune(fiveplet)), " at index ", key.index, " by five at ", index, "\n")
					if found_keys == Target {
						return h.ToString(key.index)
					}
				}
			}
		}
	}
	return "Not found"
}

func PartOne(lines []string) string {
	return Do(lines[0], 0)
}

func PartTwo(lines []string) string {
	return Do(lines[0], 2016)
}
