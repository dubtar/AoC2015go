package helpers

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func InputLines(day int) ([]string, error) {
	file, err := os.Open(fmt.Sprintf("inputs/%d", day))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not find input file for day %d. Add input to ./inputs/%d\n", day, day)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func ToInt(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	Check(err)
	return val
}

func ToString[T any](i T) string {
	return fmt.Sprint(i)
}

func IsLowChar(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func Pow(base, exponent int64) int64 {
	return int64(math.Pow(float64(base), float64(exponent)))

}

func Combinations[T any](set []T) (subsets [][]T) {
	length := uint(len(set))

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []T

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

func Multiply[T int | int64](list []T) T {
	res := T(1)
	for _, e := range list {
		res *= e
	}
	return res
}
