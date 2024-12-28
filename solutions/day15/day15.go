package day15

import (
	"fmt"
	h "go-aoc-template/internal/helpers"
	"math"
	"strings"
)

type Ingredient []int64

func NewIngredient(line string) Ingredient {
	vals := strings.Split(strings.Split(line, ": ")[1], ", ")

	var result Ingredient
	for _, v := range vals {
		result = append(result, h.ToInt(strings.Split(v, " ")[1]))
	}

	return result
}

func parse(lines []string) []Ingredient {
	ings := make([]Ingredient, 0)
	for _, line := range lines {
		ings = append(ings, NewIngredient(line))
	}

	return ings
}

func PartOne(lines []string) string {
	ings := parse(lines)
	limit := int64(100)
    target_calories := int64(500)
	best := int64(0)
	best2 := int64(0)
	var ings_len int = len(ings)
	for i := int64(0); float64(i) <= math.Pow(float64(limit), float64(ings_len-1)); i++ {
        cur := i
		scores := make([]int64, 5)
		used := int64(0)
		for j := 0; j < ings_len; j++ {
			q := cur % limit
			if j == ings_len-1 {
				q = limit - used
            }
            for k := 0; k < 5; k++ {
                scores[k] += ings[j][k] * q
            }
            used += q
            cur /= limit
		}
        score := int64(1)
        for k := 0; k < len(scores) - 1; k++ {
            score *= max(0, scores[k])
        }
        best = max(best, score)
        if scores[len(scores)-1] == target_calories {
            best2 = max(best2, score)
        }
	}
    return fmt.Sprint(best, best2)
}

func PartTwo(lines []string) string {
	return "Look before"
}
