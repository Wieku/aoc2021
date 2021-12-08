package main

import (
	"aoc2021/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var lines = util.ReadLines("day7/input.txt")

func main() {
	start := strings.Split(lines[0], ",")

	l := len(start)

	positions := make([]int, l)

	mean := 0

	for i, t := range start {
		positions[i], _ = strconv.Atoi(t)

		mean += positions[i]
	}

	mean /= l

	sort.Ints(positions)

	median := (positions[(l-1)/2] + positions[l/2-1]) / 2

	minFuel := 0

	minFuel2L := 0
	minFuel2U := 0

	for i := 0; i < l; i++ {
		bFuel := abs(median - positions[i])

		bFuel2L := abs(positions[i] - mean)
		bFuel2U := abs(positions[i] - (mean + 1))

		minFuel += bFuel

		minFuel2L += bFuel2L * (bFuel2L + 1) / 2
		minFuel2U += bFuel2U * (bFuel2U + 1) / 2
	}

	fmt.Println("Part1:", minFuel)

	if minFuel2L > minFuel2U {
		fmt.Println("Part2:", minFuel2U)
	} else {
		fmt.Println("Part2:", minFuel2L)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
