package day18

import (
	"fmt"
	"strconv"
)

// import (
//     h "go-aoc-template/internal/helpers"
// )

type Grid [][]bool

var Neighbours = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func newValue(grid Grid, i, j int) bool {
	neighboursOn := 0
	for _, neighbour := range Neighbours {
		ni := i + neighbour[0]
		nj := j + neighbour[1]
		if 0 <= ni && ni < len(grid) && 0 <= nj && nj < len(grid[0]) && grid[ni][nj] {
			neighboursOn++
		}
	}
	if grid[i][j] {
		return neighboursOn == 2 || neighboursOn == 3
	}
	return neighboursOn == 3
}

func PartOne(lines []string) string {
	N := len(lines)
	grid := make(Grid, 0, N)
	for _, line := range lines {
		row := make([]bool, len(line))
		for i, v := range line {
			row[i] = v == '#'
		}
		grid = append(grid, row)
	}
	isTest := N == 6
	steps := 100
	if isTest {
		steps = 4
	}
	for range steps {
		newgrid := make(Grid, 0, N)
		for i := 0; i < N; i++ {
			row := make([]bool, N)
			for j := 0; j < N; j++ {
				row[j] = newValue(grid, i, j)
			}
			newgrid = append(newgrid, row)
		}
		grid = newgrid
		if isTest {
			fmt.Println()
			for _, row := range grid {
				for _, v := range row {
					if v {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println()
			}
		}
	}
	result := 0
	for _, row := range grid {
		for _, v := range row {
			if v {
				result++
			}
		}
	}
	return strconv.Itoa(result)
}

func setCorners(grid Grid) {
	N := len(grid) - 1
	for _, c := range [][]int{{0, 0}, {0, N}, {N, 0}, {N, N}} {
		grid[c[0]][c[1]] = true
	}
}

func PartTwo(lines []string) string {
	N := len(lines)
	grid := make(Grid, 0, N)
	for _, line := range lines {
		row := make([]bool, len(line))
		for i, v := range line {
			row[i] = v == '#'
		}
		grid = append(grid, row)
	}
	isTest := N == 6
	steps := 100
	if isTest {
		steps = 5
	}
	setCorners(grid)
	for range steps {
		newgrid := make(Grid, 0, N)
		for i := 0; i < N; i++ {
			row := make([]bool, N)
			for j := 0; j < N; j++ {
				row[j] = newValue(grid, i, j)
			}
			newgrid = append(newgrid, row)
		}
		grid = newgrid
		setCorners(grid)
		if isTest {
			fmt.Println()
			for _, row := range grid {
				for _, v := range row {
					if v {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println()
			}
		}
	}
	result := 0
	for _, row := range grid {
		for _, v := range row {
			if v {
				result++
			}
		}
	}
	return strconv.Itoa(result)
}
