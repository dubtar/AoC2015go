package day22

import (
	"fmt"
	h "go-aoc-template/internal/helpers"
	"os"
	"strings"
)

type Cell struct {
	used      int
	total     int
	available int
	coord     h.Coords
}

func buildGrid(lines []string) (grid map[h.Coords]Cell, maxX int64, maxY int64) {
	grid = make(map[h.Coords]Cell, len(lines))
	maxX = -1
	maxY = -1

	for _, line := range lines {
		if strings.HasPrefix(line, "/dev/") {
			cell := Cell{}
			_, err := fmt.Sscanf(line, "/dev/grid/node-x%d-y%d %dT %dT %dT", &cell.coord.X,
				&cell.coord.Y, &cell.total, &cell.used, &cell.available)
			if err != nil {
				panic(err)
			}
			grid[cell.coord] = cell
			maxX = max(maxX, cell.coord.X)
			maxY = max(maxY, cell.coord.Y)
		}
	}
	return grid, maxX, maxY
}

func PartOne(lines []string) string {
	grid, _, _ := buildGrid(lines)

	result := 0
	for _, cell := range grid {
		if cell.used == 0 { // 1. cell A is not empty
			continue
		}
		for _, other := range grid {
			if cell.coord == other.coord { // 2. Nodes A and B are not the same node.
				continue
			}
			if cell.used <= other.available { // 3. The data on node A (its Used) would fit on node B (its Avail).
				result++
			}
		}
	}
	return fmt.Sprintf("%d", result)
}

func PartTwo(lines []string) string {
	grid, maxX, maxY := buildGrid(lines)

	// finish := h.Coords{X: 0, Y: 0}
	// target := h.Coords{X: maxX, Y: 0}
	zero := h.Coords{X: -1, Y: -1}
	// save to /tmp/grid.txt file
	file, err := os.Create("/tmp/grid.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for y := range maxY {
		for x := range maxX {
			fmt.Fprintf(file, "%3d/%3d  ", grid[h.Coords{X: x, Y: y}].used, grid[h.Coords{X: x, Y: y}].total)
		}
		fmt.Fprintln(file)
	}
	for c, cell := range grid {
		if cell.used == 0 {
			zero = c
			// break
		}
	}
	if zero.X < 0 {
		panic("zero cell not found")
	}

	result := zero.Y         // столько раз двигаем до первого ряда (упрощение, что 0 < zero.X < target.X)
	result += maxX - zero.X  // столько раз двигаем до целевого столбца ( теперь target.X = maxX - 1)
	result += (maxX - 1) * 5 // чтобы перейдвинуть цель на 1 влево
	if maxX > 5 {
		result += 12 // магическая добавка, глазами изучив данные увидел стенку, на обход которой нужно 6 в сторону шагов и 6 назад
	}

	return fmt.Sprintf("%d", result)
}
