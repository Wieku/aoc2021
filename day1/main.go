package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
)

var lines = util.ReadLines("day1/input.txt")

func main() {
	vals := make([]int, len(lines))
	for i, t := range lines {
		vals[i], _ = strconv.Atoi(t)
	}

	vP := vals[0]

	count := 0

	for i, v := range vals {
		if i > 0 && v > vP {
			count++
		}

		vP = v
	}

	fmt.Println(count)

	vPS := vals[0] + vals[1] + vals[2]

	countS := 0

	for i := 1; i < len(vals)-2; i++ {
		vS := vals[i] + vals[i+1] + vals[i+2]

		if vS > vPS {
			countS++
		}

		vPS = vS
	}

	fmt.Println(countS)
}
