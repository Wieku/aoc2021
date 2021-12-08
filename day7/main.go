package main

import (
	"aoc2021/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var lines = util.ReadLines("day7/input.txt")

func main() {
	start := strings.Split(lines[0], ",")

	positions := make([]int, len(start))

	mx := 0

	for i, t := range start {
		positions[i], _ = strconv.Atoi(t)

		if positions[i] > mx {
			mx = positions[i]
		}
	}

	fuel := make([]int, mx+1)
	fuel2 := make([]int, mx+1)

	for i := 0; i < len(fuel); i++ {
		for j := 0; j < len(positions); j++ {
			bFuel := abs(i - positions[j])

			fuel[i] += bFuel
			fuel2[i] += bFuel * (bFuel + 1) / 2
		}
	}

	minFuel := math.MaxInt
	minFuel2 := math.MaxInt

	for i := 0; i < len(fuel); i++ {
		if fuel[i] < minFuel {
			minFuel = fuel[i]
		}

		if fuel2[i] < minFuel2 {
			minFuel2 = fuel2[i]
		}
	}

	fmt.Println(minFuel)
	fmt.Println(minFuel2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
