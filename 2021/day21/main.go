package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"strings"
)

var lines = util.ReadLines("2021/day21/input.txt")

var ppos []int

func init() {
	for _, k := range lines {
		a, _ := strconv.Atoi(strings.Split(k, " ")[4])
		ppos = append(ppos, a)
	}
}

func main() {
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
}

func p1() int {
	positions := make([]int, len(ppos))
	scores := make([]int, len(ppos))
	copy(positions, ppos)

	var rVal, rolls, lost int

	for shallC := true; shallC; {
		for i := range positions {
			r1 := roll1(rVal)
			r2 := roll1(r1)
			r3 := roll1(r2)

			rVal = r3

			positions[i] += r1 + r2 + r3

			rolls += 3

			for positions[i] > 10 {
				positions[i] -= 10
			}

			scores[i] += positions[i]

			if scores[i] >= 1000 {
				lost = (i + 1) % 2
				shallC = false
				break
			}
		}
	}

	return rolls * scores[lost]
}

func roll1(value int) int {
	return value%100 + 1
}

func p2() int {
	// Minimum and maximum score from 3 rolls is 3 and 9 so we store the number of universes with the same dice results, massively reduces the number of recursed calls
	cache := make([]int, 7)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				cache[i+j+k-3]++ // -3 so we fit in the array
			}
		}
	}

	pw := roll2(0, [2]int{ppos[0], ppos[1]}, [2]int{0, 0}, 1, 21, cache)

	if pw[0] > pw[1] {
		return pw[0]
	}

	return pw[1]
}

func roll2(t int, pp [2]int, ps [2]int, universes, mScore int, cache []int) (pw [2]int) {
	p0 := pp[t]
	ps0 := ps[t]

	for i, k := range cache {
		pos := p0 + (i + 3) // +3 to reverse -3 from preparation stage
		for pos > 10 {
			pos -= 10
		}

		pp[t] = pos
		ps[t] = ps0 + pos

		if ps[t] >= mScore {
			pw[t] += universes * k
		} else {
			po := roll2((t+1)%2, pp, ps, universes*k, mScore, cache)
			pw[0] += po[0]
			pw[1] += po[1]
		}
	}

	return
}
