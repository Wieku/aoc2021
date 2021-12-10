package main

import (
	"aoc2021/util"
	"fmt"
	"sort"
)

var lines = util.ReadLines("day10/input.txt")

func main() {
	lut := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
		'>': '<',
	}

	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	p1 := 0

	var p2 []int

	for _, a := range lines {
		var stack []rune
		var broken bool

		for _, b := range []rune(a) {
			if r, ok := lut[b]; ok {
				if stack[len(stack)-1] != r {
					p1 += scores[b]
					broken = true
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, b)
			}
		}

		if !broken && len(stack) > 0 {
			sc := 0
			for i := len(stack) - 1; i >= 0; i-- {
				sc *= 5
				sc += scores[stack[i]]
			}

			p2 = append(p2, sc)
		}
	}

	sort.Ints(p2)

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2[len(p2)/2])
}
