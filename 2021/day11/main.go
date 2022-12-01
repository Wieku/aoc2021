package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"strings"
)

var lines = util.ReadLines("day11/input.txt")

var data [][]int

func main() {
	for _, t := range lines {
		s := strings.Split(t, "")

		n := make([]int, len(s))

		for i, k := range s {
			n[i], _ = strconv.Atoi(k)
		}

		data = append(data, n)
	}

	type vec struct {
		y, x int
	}

	flashSum := 0

	p2Solved := false

	for i := 0; i < 100 || !p2Solved; i++ {
		mark := make([][]bool, len(data))
		for j := range mark {
			mark[j] = make([]bool, len(data[0]))
		}

		var stack []vec

		for y, yd := range data {
			for x := range yd {
				data[y][x]++

				if data[y][x] > 9 {
					stack = append(stack, vec{y, x})
				}
			}
		}

		for len(stack) > 0 {
			el := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if mark[el.y][el.x] {
				continue
			}

			mark[el.y][el.x] = true

			for ny := max(0, el.y-1); ny <= min(el.y+1, len(data)-1); ny++ {
				for nx := max(0, el.x-1); nx <= min(el.x+1, len(data[0])-1); nx++ {
					if !(nx == el.x && ny == el.y) {
						data[ny][nx]++

						if data[ny][nx] > 9 && !mark[ny][nx] {
							stack = append(stack, vec{ny, nx})
						}
					}
				}
			}
		}

		cFlash := 0

		for y, yd := range data {
			for x := range yd {
				if data[y][x] > 9 {
					cFlash++
					data[y][x] = 0
				}
			}
		}

		flashSum += cFlash

		if i == 99 {
			fmt.Println("Part 1:", flashSum)
		}

		if !p2Solved && cFlash == len(data)*len(data[0]) {
			fmt.Println("Part 2:", i+1)
			p2Solved = true
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
