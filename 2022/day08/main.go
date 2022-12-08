package main

import (
	"aoc2021/util"
	"fmt"
)

var lines = util.ReadLines("2022/day08/input.txt")

var grid [][]int8

func init() {
	for _, line := range lines {
		var row []int8
		for i := 0; i < len(line); i++ {
			row = append(row, int8(line[i]-'0'))
		}

		grid = append(grid, row)
	}

}

func main() {
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
}

func p1() (sum int) {
	mark := make([][]int8, len(grid))
	for i, row := range grid {
		mark[i] = make([]int8, len(row))
	}

	for y := 0; y < len(grid); y++ {
		rowMax1 := int8(-1)
		rowMax2 := int8(-1)

		for x := range grid[y] {
			x1 := x
			x2 := len(grid[y]) - 1 - x

			cell1 := grid[y][x1]
			cell2 := grid[y][x2]

			if cell1 > rowMax1 {
				mark[y][x1] = 1

				rowMax1 = cell1
			}

			if cell2 > rowMax2 {
				mark[y][x2] = 1

				rowMax2 = cell2
			}
		}

	}

	for x := 0; x < len(grid[0]); x++ {
		rowMax1 := int8(-1)
		rowMax2 := int8(-1)

		for y := range grid {
			y1 := y
			y2 := len(grid) - 1 - y

			cell1 := grid[y1][x]
			cell2 := grid[y2][x]

			if cell1 > rowMax1 {
				mark[y1][x] = 1

				rowMax1 = cell1
			}

			if cell2 > rowMax2 {
				mark[y2][x] = 1

				rowMax2 = cell2
			}
		}

	}

	for _, row := range mark {
		for _, cell := range row {
			sum += int(cell)
		}
	}

	return
}

func p2() (max int) {
	height := len(grid)
	for y := 1; y < height-1; y++ {
		width := len(grid[y])

		for x := 1; x < width-1; x++ {
			cell := grid[y][x]

			var left, right, top, bottom int

			for sX := 1; sX <= x; sX++ {
				left++

				if grid[y][x-sX] >= cell {
					break
				}
			}

			for sX := 1; sX <= width-x-1; sX++ {
				right++

				if grid[y][x+sX] >= cell {
					break
				}
			}

			for sY := 1; sY <= y; sY++ {
				top++

				if grid[y-sY][x] >= cell {
					break
				}
			}

			for sY := 1; sY <= height-y-1; sY++ {
				bottom++

				if grid[y+sY][x] >= cell {
					break
				}
			}

			mult := left * right * top * bottom

			if mult > max {
				max = mult
			}
		}
	}

	return
}
