package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"strings"
)

var lines = util.ReadLines("day13/input.txt")

type vec struct {
	x, y int
}

var paper = make(map[vec]uint8)

var folds []string

func main() {
	for _, t := range lines {
		if strings.HasPrefix(t, "fold along") {
			folds = append(folds, strings.TrimPrefix(t, "fold along "))
			continue
		}

		if s := strings.Split(t, ","); len(s) == 2 {
			x, _ := strconv.Atoi(s[0])
			y, _ := strconv.Atoi(s[1])

			paper[vec{x, y}] = 1
		}
	}

	for i, a := range folds {
		f := strings.Split(a, "=")

		o, _ := strconv.Atoi(f[1])

		for k := range paper {
			if f[0] == "x" && k.x > o {
				paper[vec{2*o - k.x, k.y}] = 1
				delete(paper, k)
			} else if f[0] == "y" && k.y > o {
				paper[vec{k.x, 2*o - k.y}] = 1
				delete(paper, k)
			}
		}

		fmt.Println("Fold", i, len(paper))
	}

	var mX, mY int

	for k := range paper {
		if k.x > mX {
			mX = k.x
		}

		if k.y > mY {
			mY = k.y
		}
	}

	for y := 0; y <= mY; y++ {
		for x := 0; x <= mX; x++ {
			if _, ok := paper[vec{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
