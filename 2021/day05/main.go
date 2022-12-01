package main

import (
	"aoc2021/util"
	"fmt"
	"regexp"
	"strconv"
)

type vec struct {
	x, y int
}

var board = make(map[vec]int)
var board2 = make(map[vec]int)

var lines = util.ReadLines("day05/input.txt")

func main() {
	r, err := regexp.Compile("(,| -> )")
	if err != nil {
		panic(err)
	}

	for _, t := range lines {
		s := r.Split(t, -1)

		x0, _ := strconv.Atoi(s[0])
		y0, _ := strconv.Atoi(s[1])
		x1, _ := strconv.Atoi(s[2])
		y1, _ := strconv.Atoi(s[3])

		xStep, xL := getP(x0, x1)
		yStep, yL := getP(y0, y1)

		h := xL
		if yL > xL {
			h = yL
		}

		for i := 0; i <= h; i++ {
			v := vec{x0 + xStep*i, y0 + yStep*i}

			if y0 == y1 || x0 == x1 {
				board[v] = board[v] + 1
			}

			board2[v] = board2[v] + 1
		}
	}

	sum1 := 0
	sum2 := 0

	for _, v := range board {
		if v >= 2 {
			sum1++
		}
	}

	for _, v := range board2 {
		if v >= 2 {
			sum2++
		}
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}

func getP(v0, v1 int) (step int, d int) {
	if v1 > v0 {
		step = 1
		d = v1 - v0
	} else if v1 < v0 {
		step = -1
		d = v0 - v1
	}

	return
}
