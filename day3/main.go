package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
)

var lines = util.ReadLines("day3/input.txt")

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
	textO := make([]string, len(lines))
	textS := make([]string, len(lines))

	copy(textO, lines)
	copy(textS, lines)

	siz := len(lines[0])

	for i := 0; i < siz && len(textO) > 1; i++ {
		sieve := getSieve(textO, i)

		for j := 0; j < len(textO) && len(textO) > 1; j++ {
			if textO[j][i] != sieve {
				textO = append(textO[:j], textO[j+1:]...)
				j--
			}
		}
	}

	for i := 0; i < siz && len(textS) > 1; i++ {
		sieve := getSieve(textS, i)

		for j := 0; j < len(textS) && len(textS) > 1; j++ {
			if textS[j][i] == sieve {
				textS = append(textS[:j], textS[j+1:]...)
				j--
			}
		}
	}

	g, _ := strconv.ParseInt(textO[0], 2, 64)
	s, _ := strconv.ParseInt(textS[0], 2, 64)

	fmt.Println(g * s)
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
