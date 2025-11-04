package day19

import (
	"fmt"
	"slices"
	"strconv"
)

func PartOne(lines []string) string {
	num, _ := strconv.Atoi(lines[0])
	arr := make([]int, num)
	for i := range num {
		arr[i] = 1
	}
	cur := 0
	i := cur
	isStealing := true
	for num > 1 {
		i++
		if i == len(arr) {
			i = 0
		}
		if arr[i] == 0 {
			continue
		}
		if isStealing {
			// забираем
			arr[cur] += arr[i]
			arr[i] = 0
			num--
			isStealing = false
		} else {
			cur = i
			isStealing = true
		}
	}

	return strconv.Itoa(cur + 1)
}

type chainedArray[K any] struct {
	partCapacity int
	parts [][]K
}

func NewChainedArray[K any](partCapacity int) chainedArray[K] {
	return chainedArray[K]{
		partCapacity: partCapacity,
		parts:  make([][]K, 0, 100),
	}
}
func (ca *chainedArray[K]) Get(index int) K {
	for _, part :=  range ca.parts {
		if len(part) > index {
			return part[index]
		}
		index -= len(part)
	}
	panic("Index is out of bounds")
}

func (ca *chainedArray[K]) Append(value K) {
	if len(ca.parts) > 0 && len(ca.parts[len(ca.parts)-1]) < ca.partCapacity {
		ca.parts[len(ca.parts)-1] = append(ca.parts[len(ca.parts)-1], value)
		return
	}
	newPart := make([]K, 1, ca.partCapacity)
	newPart[0] = value
	ca.parts = append(ca.parts, newPart)
}

func (ca *chainedArray[K]) Delete(index int) {
	for i, part :=  range ca.parts {
		if len(part) > index {
			part = slices.Delete(part, index, index+1)
			if len(part) == 0 {
				ca.parts = slices.Delete(ca.parts, i, i+1)
			} else {
				ca.parts[i] = part
			}
			return
		}
		index -= len(part)
	}
	panic("Index is out of bounds")
}

func PartTwo(lines []string) string {
	num, _ := strconv.Atoi(lines[0])
	t := 100
	if num < 100 {
		t = 2
	}
	arr := NewChainedArray[int](t)
	for i := range num {
		arr.Append(i + 1)
	}
	cur := 0
	for num > 1 {
		target := (cur + num / 2) % num
		arr.Delete(target)
		num--
		fmt.Printf("\r n=%v", num)
		if target > cur {
			cur++
		}
		cur = cur % num
	}

	return strconv.Itoa(arr.Get(0))
}
