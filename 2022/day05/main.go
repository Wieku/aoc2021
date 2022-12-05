package main

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

var lines = util.ReadLines("2022/day05/input.txt")

func main() {
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
}

func p1() string {
	return process(true)
}

func p2() string {
	return process(false)
}

func process(reverse bool) string {
	var stacks [][]uint8

	caught := false

	for i, line := range lines {
		if !caught && line[1] == '1' {

			for j := 1; j < len(line); j += 4 {
				var stack []uint8

				for k := i - 1; k >= 0; k-- {
					if len(lines[k]) <= j || lines[k][j] == ' ' {
						break
					}

					stack = append(stack, lines[k][j])
				}

				stacks = append(stacks, stack)
			}

			caught = true
			continue
		}

		if caught && line != "" {
			spl := strings.Split(line, " ")

			amount := util.Atoi(spl[1])
			from := util.Atoi(spl[3])
			to := util.Atoi(spl[5])

			cstack := stacks[from-1]

			arr := cstack[len(cstack)-amount:]

			stacks[from-1] = cstack[:len(cstack)-amount]

			if reverse {
				util.Reverse(arr)
			}

			stacks[to-1] = append(stacks[to-1], arr...)
		}
	}

	cS := ""

	for i := 0; i < len(stacks); i++ {
		cS += string(stacks[i][len(stacks[i])-1])
	}

	return cS
}
