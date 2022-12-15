package main

import (
	"aoc2021/util"
	"fmt"
	"log"
	"strings"
)

var lines = util.ReadLines("2022/day14/input.txt")

type vec struct {
	x, y int
}

func main() {
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
}

func p1() int {
	cave := make(map[vec]uint8)

	maxY := -1

	for _, line := range lines {
		spl := strings.Split(line, " -> ")

		x0 := 0
		y0 := 0

		for i := 0; i < len(spl); i++ {
			spl2 := strings.Split(spl[i], ",")

			x1 := util.Atoi(spl2[0])
			y1 := util.Atoi(spl2[1])

			if i > 0 {
				x3 := util.Min(x0, x1)
				y3 := util.Min(y0, y1)

				x4 := util.Max(x0, x1)
				y4 := util.Max(y0, y1)

				for x := x3; x <= x4; x++ {
					for y := y3; y <= y4; y++ {
						maxY = util.Max(maxY, y)
						cave[vec{x, y}] = 1
					}
				}

			}

			x0 = x1
			y0 = y1
		}
	}

	i := 0

	escaped := false

	log.Println(maxY)

	for ; !escaped; i++ {
		x, y := 500, 0

		for {
			if y >= maxY {
				escaped = true
				break
			} else if cave[vec{x, y + 1}] == 0 {
				y++
			} else if cave[vec{x - 1, y + 1}] == 0 {
				x--
				y++
			} else if cave[vec{x + 1, y + 1}] == 0 {
				x++
				y++
			} else {
				cave[vec{x, y}] = 1
				break
			}
		}
	}

	return i - 1
}

func p2() int {
	cave := make(map[vec]uint8)

	maxY := -1

	for _, line := range lines {
		spl := strings.Split(line, " -> ")

		x0 := 0
		y0 := 0

		for i := 0; i < len(spl); i++ {
			spl2 := strings.Split(spl[i], ",")

			x1 := util.Atoi(spl2[0])
			y1 := util.Atoi(spl2[1])

			if i > 0 {
				x3 := util.Min(x0, x1)
				y3 := util.Min(y0, y1)

				x4 := util.Max(x0, x1)
				y4 := util.Max(y0, y1)

				for x := x3; x <= x4; x++ {
					for y := y3; y <= y4; y++ {
						maxY = util.Max(maxY, y)
						cave[vec{x, y}] = 1
					}
				}

			}

			x0 = x1
			y0 = y1
		}
	}

	i := 0

	log.Println(maxY)

	for ; ; i++ {
		x, y := 500, 0

		for {
			if y >= maxY+1 {
				cave[vec{x, y}] = 1
				break
			} else if cave[vec{x, y + 1}] == 0 {
				y++
			} else if cave[vec{x - 1, y + 1}] == 0 {
				x--
				y++
			} else if cave[vec{x + 1, y + 1}] == 0 {
				x++
				y++
			} else {
				cave[vec{x, y}] = 1
				break
			}
		}

		if x == 500 && y == 0 {
			break
		}
	}

	return i + 1
}
