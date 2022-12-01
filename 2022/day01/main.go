package main

import (
	"aoc2021/util"
	"fmt"
	"golang.org/x/exp/slices"
)

var lines = util.ReadLines("2022/day01/input.txt")

func main() {
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
}

func p1() int {
	var sum, biggest int

	for _, l := range lines {
		if l == "" {
			biggest = util.Max(biggest, sum)

			sum = 0

			continue
		}

		sum += util.Atoi(l)
	}

	return biggest
}

func p2() int {
	var arr []int
	var sum int

	for _, l := range lines {
		if l == "" {
			arr = append(arr, sum)

			sum = 0

			continue
		}

		sum += util.Atoi(l)
	}

	slices.SortFunc(arr, func(a, b int) bool {
		return a > b
	})

	return arr[0] + arr[1] + arr[2]
}
