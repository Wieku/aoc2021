package main

import (
	"aoc2021/util"
	"fmt"
)

var lines = util.ReadLines("2022/day02/input.txt")

var scoreMap = map[uint8]int{
	'A': 1,
	'B': 2,
	'C': 3,
	'X': 1,
	'Y': 2,
	'Z': 3,
}

func main() {
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
}

func p1() int {
	score := 0

	for _, l := range lines {
		eS := scoreMap[l[0]]
		uS := scoreMap[l[2]]

		if eS == uS {
			score += 3
		} else if (eS == 3 && uS == 1) || (eS == 1 && uS == 2) || (eS == 2 && uS == 3) {
			score += 6
		}

		score += uS
	}

	return score
}

func p2() int {
	score := 0

	for _, l := range lines {
		eS := scoreMap[l[0]]

		switch l[2] {
		case 'X':
			if eS == 1 {
				score += 3
			} else {
				score += eS - 1
			}
		case 'Y':
			score += 3 + eS
		case 'Z':
			score += 6

			if eS == 3 {
				score += 1
			} else {
				score += eS + 1
			}
		}
	}

	return score
}
