package day16

// import (
//     h "go-aoc-template/internal/helpers"
// )

func PartOne(lines []string) string {
	Target := 272
	input := lines[0]
	if len(input) < 15 {
		Target = 20 // test
	}
	return Run(input, Target)
}
func Run(input string, Target int) string {
	data := make([]byte, len(input), Target)
	for i := 0; i < len(input); i++ {
		data[i] = input[i]
	}
	for len(data) < Target {
		prevLen := len(data)
		data = append(data, '0')
		for i := 0; i < prevLen; i++ {
			v := data[prevLen-i-1]
			if v == '1' {
				data = append(data, '0')
			} else {
				data = append(data, '1')
			}
			if len(data) >= Target {
				break
			}
		}
	}
	// for _, v := range data {
	// 	print(string(v))
	// }
	// print('\n')
	checksum := buildCheckSum(data)
	for len(checksum)%2 == 0 {
		checksum = buildCheckSum(checksum)
	}
	return string(checksum)
}

func buildCheckSum(data []byte) []byte {
	result := make([]byte, len(data)/2)
	for i := 0; i < len(data); i += 2 {
		if data[i] == data[i+1] {
			result[i/2] = '1'
		} else {
			result[i/2] = '0'
		}
	}
	return result
}

func PartTwo(lines []string) string {
	target := 35651584
	return Run(lines[0], target)
}
