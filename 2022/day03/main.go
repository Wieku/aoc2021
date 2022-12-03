package main

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

var lines = util.ReadLines("2022/day03/input.txt")

func main() {
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
}

func p1() int {
	sum := 0

	for _, line := range lines {
		ln := len(line) / 2

		left, right := line[:ln], line[ln:]

		for j := 0; j < len(left); j++ {
			cChar := left[j]

			if strings.IndexByte(right, cChar) > -1 {
				sum += getScore(cChar)
				break
			}
		}
	}

	return sum
}

func p2() int {
	sum := 0

	for i := 0; i < len(lines); i += 3 {
		first, second, third := lines[i], lines[i+1], lines[i+2]

		for j := 0; j < len(first); j++ {
			cChar := first[j]

			if strings.IndexByte(second, cChar) > -1 && strings.IndexByte(third, cChar) > -1 {
				sum += getScore(cChar)
				break
			}
		}
	}

	return sum
}

func getScore(k uint8) int {
	if k >= 'a' {
		return int(k - 'a' + 1)
	}

	return int(k - 'A' + 27)
}
