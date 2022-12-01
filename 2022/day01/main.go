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
	return solve()[0]
}

func p2() int {
	arr := solve()
	return arr[0] + arr[1] + arr[2]
}

func solve() []int {
	arr := []int{0}

	for _, l := range lines {
		if l == "" {
			arr = append(arr, 0)

			continue
		}

		arr[len(arr)-1] += util.Atoi(l)
	}

	slices.SortFunc(arr, func(a, b int) bool {
		return a > b
	})

	return arr
}
