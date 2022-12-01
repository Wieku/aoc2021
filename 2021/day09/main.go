package main

import (
	"aoc2021/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var lines = util.ReadLines("day09/input.txt")

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

	p1()
	p2()
}

func p1() {
	sum := 0

	for i, a := range data {
		for j, b := range a {
			if b == 9 {
				continue
			}

			lower := true

			if (j > 0 && a[j-1] < b) ||
				(j < len(a)-1 && a[j+1] < b) ||
				(i > 0 && data[i-1][j] < b) ||
				(i < len(data)-1 && data[i+1][j] < b) {
				lower = false
			}

			if lower {
				sum += b + 1
			}
		}
	}

	fmt.Println("Part 1:", sum)
}

func p2() {
	mark := make([][]bool, len(data))
	for i := range mark {
		mark[i] = make([]bool, len(data[0]))
	}

	var basins []int

	for y, a := range data {
		for x, b := range a {
			if b == 9 || mark[y][x] {
				continue
			}

			basins = append(basins, step(data, mark, y, x))
		}
	}

	sort.Ints(basins)

	result := basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]

	fmt.Println("Part 2:", result)
}

func step(data [][]int, mark [][]bool, y, x int) (sum int) {
	mark[y][x] = true

	if y > 0 && data[y-1][x] < 9 && !mark[y-1][x] {
		sum += step(data, mark, y-1, x)
	}

	if y < len(data)-1 && data[y+1][x] < 9 && !mark[y+1][x] {
		sum += step(data, mark, y+1, x)
	}

	if x > 0 && data[y][x-1] < 9 && !mark[y][x-1] {
		sum += step(data, mark, y, x-1)
	}

	if x < len(data[y])-1 && data[y][x+1] < 9 && !mark[y][x+1] {
		sum += step(data, mark, y, x+1)
	}

	return sum + 1
}
