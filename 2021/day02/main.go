package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"strings"
)

var lines = util.ReadLines("day02/input.txt")

func main() {
	horiz := 0
	depth := 0

	aim := 0
	depth2 := 0

	for _, t := range lines {
		s := strings.Split(t, " ")

		v, _ := strconv.Atoi(s[1])

		switch s[0] {
		case "down":
			depth += v
			aim += v
		case "up":
			depth -= v
			aim -= v
		case "forward":
			depth2 += aim * v
			horiz += v
		}
	}

	fmt.Println(horiz * depth)
	fmt.Println(horiz * depth2)
}
