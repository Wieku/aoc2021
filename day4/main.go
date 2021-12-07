package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"strings"
)

const bSize = 5

type field struct {
	val     string
	checked bool
}

type board [bSize * bSize]field

var lines = util.ReadLines("day4/input.txt")

func main() {
	var lastWon *board
	var lastRnd string

	wonOnce := false

	var boards []*board

	for i := 2; i < len(lines); i += 6 {
		var bd board

		for j := 0; j < bSize; j++ {
			for k, v := range strings.Fields(lines[i+j]) {
				bd[j*bSize+k] = field{val: v}
			}
		}

		boards = append(boards, &bd)
	}

	rnds := strings.Split(lines[0], ",")

	for _, v := range rnds {
		for i := 0; i < len(boards); i++ {
			bd := boards[i]

			check(bd, v)

			if isBingo(bd) {
				if !wonOnce {
					printScore(bd, v)
					wonOnce = true
				}

				lastWon = bd
				lastRnd = v

				boards = append(boards[:i], boards[i+1:]...)
				i--
			}
		}
	}

	if lastWon != nil {
		printScore(lastWon, lastRnd)
	}

}

func printScore(bd *board, rnd string) {
	sum := 0

	for i := 0; i < bSize*bSize; i++ {
		if !bd[i].checked {
			v, _ := strconv.Atoi(bd[i].val)
			sum += v
		}
	}

	r, _ := strconv.Atoi(rnd)

	fmt.Println(r * sum)
}

func check(bd *board, s string) {
	for i := 0; i < bSize*bSize; i++ {
		if bd[i].val == s {
			bd[i].checked = true
		}
	}
}

func isBingo(bd *board) bool {
	for i := 0; i < bSize; i++ {
		horiz := true
		vert := true

		for j := 0; j < bSize; j++ {
			if !bd[i*bSize+j].checked {
				horiz = false
			}

			if !bd[j*bSize+i].checked {
				vert = false
			}
		}

		if horiz || vert {
			return true
		}
	}

	return false
}
