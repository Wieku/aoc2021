package main

import (
	"aoc2021/util"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type vec struct {
	y, x int
}

var lines = util.ReadLines("day15/input.txt")

var data [][]int

func init() {
	for _, t := range lines {
		s := strings.Split(t, "")

		n := make([]int, len(s))

		for i, k := range s {
			n[i], _ = strconv.Atoi(k)
		}

		data = append(data, n)
	}
}

// VERY VERY SLOW BUT WORKS, DON'T JUDGE

func main() {
	fmt.Println("Part 1:", traverse(data))

	dY := len(data)
	dX := len(data[0])

	data2 := make([][]int, dY*5)
	for y := range data2 {
		data2[y] = make([]int, dX*5)

		for x := range data2[y] {
			data2[y][x] = data[y%dY][x%dX] + y/dY + x/dX
			if data2[y][x] > 9 {
				data2[y][x] -= 9
			}
		}
	}

	fmt.Println("Part 2:", traverse(data2))
}

func traverse(graph [][]int) int {
	dY := len(graph)
	dX := len(graph[0])

	dists := make([][]int, dY)
	for y := range dists {
		dists[y] = make([]int, dX)
		for x := range dists[y] {
			dists[y][x] = math.MaxInt
		}
	}

	dists[0][0] = 0

	stack := []vec{{0, 0}}

	for len(stack) > 0 {
		el := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		nds := []vec{{el.y - 1, el.x}, {el.y + 1, el.x}, {el.y, el.x - 1}, {el.y, el.x + 1}}

		for _, n := range nds {
			if !(n.y >= 0 && n.y <= dY-1 &&
				n.x >= 0 && n.x <= dX-1) {
				continue
			}

			cD := dists[el.y][el.x] + graph[n.y][n.x]

			if cD < dists[n.y][n.x] {
				dists[n.y][n.x] = cD
				stack = append(stack, n)
			}
		}

		sort.Slice(stack, func(i, j int) bool {
			v1 := stack[i]
			v2 := stack[j]
			return dists[v1.y][v1.x] > dists[v2.y][v2.x]
		})
	}

	return dists[dY-1][dX-1]
}
