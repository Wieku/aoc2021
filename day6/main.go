package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"strings"
)

var lines = util.ReadLines("day6/input.txt")

func main() {
	start := strings.Split(lines[0], ",")

	var bins [9]int

	for _, t := range start {
		bin, _ := strconv.Atoi(t)
		bins[bin]++
	}

	for i := 1; i <= 256; i++ {
		sp := bins[0]

		for j := 1; j <= 8; j++ {
			bins[j-1] = bins[j]
		}

		bins[6] += sp
		bins[8] = sp

		if i == 80 {
			sum := 0
			for _, v := range bins {
				sum += v
			}

			fmt.Println(sum)
		}
	}

	sum := 0
	for _, v := range bins {
		sum += v
	}

	fmt.Println(sum)
}
