package main

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

var lines = util.ReadLines("2022/day04/input.txt")

type pair struct {
	firstLower int
	firstUpper int

	secondLower int
	secondUpper int
}

var pairs []*pair

func init() {
	for _, line := range lines {
		spl := strings.Split(line, ",")
		first := strings.Split(spl[0], "-")
		second := strings.Split(spl[1], "-")

		pairs = append(pairs, &pair{
			firstLower:  util.Atoi(first[0]),
			firstUpper:  util.Atoi(first[1]),
			secondLower: util.Atoi(second[0]),
			secondUpper: util.Atoi(second[1]),
		})
	}
}

func main() {
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
}

func p1() int {
	sum := 0

	for _, p := range pairs {
		if (p.secondLower <= p.firstLower && p.firstUpper <= p.secondUpper) || (p.firstLower <= p.secondLower && p.secondUpper <= p.firstUpper) {
			sum++
		}
	}

	return sum
}

func p2() int {
	sum := 0

	for _, p := range pairs {
		if p.firstLower <= p.secondUpper && p.firstUpper >= p.secondLower {
			sum++
		}
	}

	return sum
}
