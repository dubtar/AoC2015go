package helpers

import "slices"

type chainedArray[K any] struct {
	partCapacity int
	parts        [][]K
}

func NewChainedArray[K any](partCapacity int) chainedArray[K] {
	return chainedArray[K]{
		partCapacity: partCapacity,
		parts:        make([][]K, 0, 100),
	}
}
func (ca *chainedArray[K]) Get(index int) K {
	for _, part := range ca.parts {
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
	for i, part := range ca.parts {
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
