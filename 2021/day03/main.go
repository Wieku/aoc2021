package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
)

var lines = util.ReadLines("2021/day03/input.txt")

func main() {
	p1()
	p2()
}

func p1() {
	siz := len(lines[0])

	gamma := 0
	epsilon := 0

	for i := 0; i < siz; i++ {
		sm := 0

		for j := 0; j < len(lines); j++ {
			sm += int(lines[j][i]) - 48
		}

		if sm > len(lines)/2 {
			gamma |= 1 << (siz - 1 - i)
		} else {
			epsilon |= 1 << (siz - 1 - i)
		}
	}

	fmt.Println(gamma * epsilon)
}

func p2() {
	g, _ := strconv.ParseInt(getLine(lines, false), 2, 64)
	s, _ := strconv.ParseInt(getLine(lines, true), 2, 64)

	fmt.Println(g * s)
}

func getLine(src []string, inverse bool) string {
	vals := make([]string, len(src))
	copy(vals, src)

	siz := len(vals[0])

	for i := 0; i < siz && len(vals) > 1; i++ {
		sieve := getSieve(vals, i)

		for j := 0; j < len(vals) && len(vals) > 1; j++ {
			if (vals[j][i] == sieve) != inverse {
				vals = append(vals[:j], vals[j+1:]...)
				j--
			}
		}
	}

	return vals[0]
}

func getSieve(vals []string, x int) uint8 {
	sm := 0
	for i := 0; i < len(vals); i++ {
		sm += int(vals[i][x]) - 48
	}

	if sm >= (len(vals)+1)/2 {
		return '1'
	}

	return '0'
}
