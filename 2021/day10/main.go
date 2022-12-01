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

	scores := map[rune]int64{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	var (
		p1 int64
		p2 []int64
	)

	for _, a := range lines {
		var (
			stack  []rune
			broken bool
		)

		for _, b := range a {
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

		if sLen := len(stack); !broken && sLen > 0 {
			var sc int64

			for i := range stack {
				sc = sc*5 + scores[stack[sLen-1-i]]
			}

			p2 = append(p2, sc)
		}
	}

	sort.Slice(p2, func(i, j int) bool { return p2[i] < p2[j] })

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2[len(p2)/2])
}
