package day20

import (
	"fmt"
	h "go-aoc-template/internal/helpers"
)

var primes = []int64{2, 3, 5, 7, 11, 13}

func PartOne(lines []string) string {
	target := h.ToInt(lines[0]) / 10
	// fmt.Println("Primes", primes)
	progress := int64(100)
	maxs := int64(0)
	maxi := maxs
	for i := primes[len(primes)-1] + 1; ; i++ {
		ds := sumOfDividers(i, primes, false)
		if ds > target {
			return fmt.Sprint(i)
		}
		if ds == 1+i {
			primes = append(primes, i)
		}
		if i < 20 {
			fmt.Println(i, ":", ds)
		}
		if ds > maxs {
			maxs = ds
			maxi = i
		}
		if i%progress == 0 {
			progress *= 10
			fmt.Println(maxs, "at", maxi)
		}
	}
}

func sumOfDividers(i int64, primes []int64, limit50 bool) int64 {
	dels := []int64{}
	ii := i
	for _, p := range primes {
		for ii%p == 0 {
			dels = append(dels, p)
			ii /= p
		}
		if ii == 1 {
			break
		}
	}
	if ii > 1 {
		dels = append(dels, ii)
	}
	visited := map[int64]any{}
	sum := int64(1)
	for _, comb := range h.Combinations(dels) {
		mult := h.Multiply(comb)
		if limit50 && mult < i/50 {
			continue
		}
		if _, was := visited[mult]; !was {
			visited[mult] = struct{}{}
			sum += mult
		}
	}
	return sum
}

func PartTwo(lines []string) string {
	target := h.ToInt(lines[0]) / 11

	maxs := int64(0)
	maxi := maxs
	for i := primes[len(primes)-1] + 1; ; i++ {
		ds := sumOfDividers(i, primes, true)
		if ds > target {
			return fmt.Sprint(i)
		}
		if i < 20 {
			fmt.Println(i, ":", ds)
		}
		if ds > maxs {
			maxs = ds
			maxi = i
		}
		if i%100000 == 0 {
			fmt.Println(maxs, "at", maxi)
		}
	}
}

func sumOfDiviersNaive(i int64) int64 { // 5 mins
	sum := int64(1)
	for d := int64(2); d <= i/2; d++ {
		if i%d == 0 {
			sum += d
		}
	}
	return sum + i
}

func sumOfDividers2Naive(i int64, primes []int64) int64 { // 5 mins also
	dels := []int64{}
	for _, p := range primes {
		for i%p == 0 {
			dels = append(dels, p)
			i /= p
		}
		if i == 1 {
			break
		}
	}
	if i > 1 {
		dels = append(dels, i)
	}
	visited := map[int64]any{}
	sum := int64(1)
	for _, comb := range h.Combinations(dels) {
		mult := h.Multiply(comb)
		if _, was := visited[mult]; !was {
			visited[mult] = struct{}{}
			sum += mult
		}
	}
	return sum
}
