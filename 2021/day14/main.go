package main

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

var lines = util.ReadLines("2021/day14/input.txt")

var startSequence string
var substitutions = make(map[string]string)

func main() {
	for i, t := range lines {
		if i == 0 {
			startSequence = t
		} else if i > 1 {
			if s := strings.Split(t, " -> "); len(s) == 2 {
				substitutions[s[0]] = s[1]
			}
		}
	}

	counts := make(map[string]int64)
	sequence := make(map[string]int64)

	for i, k := range startSequence {
		counts[string(k)]++

		if i+1 < len(startSequence) {
			sequence[startSequence[i:i+2]]++
		}
	}

	for i := 0; i < 40; i++ {
		sequence2 := make(map[string]int64)

		for k, v := range sequence {
			if s, ok := substitutions[k]; ok {
				counts[s] += v

				sequence2[k[:1]+s] += v
				sequence2[s+k[1:]] += v
			}
		}

		sequence = sequence2

		if i == 9 {
			fmt.Println("Part 1:", getDiff(counts))
		}
	}

	fmt.Println("Part 2:", getDiff(counts))
}

func getDiff(counts map[string]int64) int64 {
	var min, max int64

	for _, v := range counts {
		if v > max {
			max = v
		}

		if min == 0 || v < min {
			min = v
		}
	}

	return max - min
}
