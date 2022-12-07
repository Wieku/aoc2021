package main

import (
	"aoc2021/util"
	"fmt"
)

var lines = util.ReadLines("2022/day06/input.txt")

var rawData []uint8
var rawData2 []uint

func init() {
	rawData = make([]uint8, len(lines[0]))
	rawData2 = make([]uint, len(lines[0]))

	for i := 0; i < len(lines[0]); i++ {
		rawData[i] = lines[0][i] - 'a'
		rawData2[i] = 1 << rawData[i]
	}
}

func main() {
	fmt.Println("Normal:")
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
	fmt.Println()
	fmt.Println("Array:")
	fmt.Println("Part 1:", p1A())
	fmt.Println("Part 2:", p2A())
	fmt.Println()
	fmt.Println("Array 2:")
	fmt.Println("Part 1:", p1A2())
	fmt.Println("Part 2:", p2A2())
	fmt.Println()
	fmt.Println("Bits:")
	fmt.Println("Part 1:", p1B())
	fmt.Println("Part 2:", p2B())
}
