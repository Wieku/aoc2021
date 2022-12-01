package main

import (
	"aoc2021/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var lines = util.ReadLines("2021/day18/input.txt")

type leaf struct {
	depth int
	right bool
	value int
	jump  int
}

var numbers [][]leaf

func init() {
	for _, k := range lines {
		numbers = append(numbers, parseNumber(k))
	}
}

func parseNumber(inp string) (list []leaf) {
	lastDepth := -1
	currDepth := -1

	for i := 0; i < len(inp); i++ {
		if inp[i] == '[' {
			currDepth++
			continue
		}

		if inp[i] == ']' {
			currDepth--
			continue
		}

		if inp[i] == ',' {
			if currDepth < lastDepth {
				list[len(list)-1].jump = abs(currDepth - lastDepth)
			}

			lastDepth = currDepth
			continue
		}

		iC := strings.IndexRune(inp[i:], ',')
		iB := strings.IndexRune(inp[i:], ']')

		if iC == -1 || iB < iC {
			vl, _ := strconv.Atoi(inp[i : i+iB])

			list = append(list, leaf{
				depth: currDepth,
				right: true,
				value: vl,
			})
		} else {
			vl, _ := strconv.Atoi(inp[i : i+iC])

			list = append(list, leaf{
				depth: currDepth,
				right: false,
				value: vl,
				jump:  abs(currDepth - lastDepth),
			})
		}

		lastDepth = currDepth
	}

	if currDepth < lastDepth {
		list[len(list)-1].jump = abs(currDepth - lastDepth)
	}

	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func main() {
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
}

func p1() int {
	sum := numbers[0]
	for _, k := range numbers[1:] {
		sum = add(sum, k)
	}

	return magnitude(sum)
}

func p2() (maxMag int) {
	for i, k := range numbers[:len(numbers)-1] {
		for _, l := range numbers[i+1:] {
			mag := magnitude(add(k, l))

			if mag > maxMag {
				maxMag = mag
			}
		}
	}

	return
}

func add(in1, in2 []leaf) (out []leaf) {
	inA := make([]leaf, 0, len(in1)+len(in2))

	for _, k := range in1 {
		k.depth++
		inA = append(inA, k)
	}

	for _, k := range in2 {
		k.depth++
		inA = append(inA, k)
	}

	inA[0].jump++
	inA[len(inA)-1].jump++

	return reduce(inA)
}

func reduce(in []leaf) (out []leaf) {
	out = make([]leaf, len(in))
	for i, k := range in {
		out[i] = k
	}

	for processMore := true; processMore; {
		processMore = false

		// explode
		for moreExplodes := true; moreExplodes; {
			moreExplodes = false

			for i := 0; i < len(out); i++ {
				l := out[i]

				if l.depth >= 4 && !l.right {
					r := out[i+1]
					if l.depth == r.depth && r.right {
						nL := leaf{
							depth: l.depth - 1,
						}

						if i-1 >= 0 {
							out[i-1].value += l.value
						}

						if i+2 < len(out) {
							out[i+2].value += r.value
						}

						if r.jump > l.jump {
							nL.right = true
							nL.jump = r.jump - 1
						} else {
							nL.jump = l.jump - 1
						}

						copy(out[i+1:], out[i+2:])

						out[i] = nL

						out = out[:len(out)-1]

						moreExplodes = true
						break
					}
				}
			}
		}

		// split
		for i := 0; i < len(out); i++ {
			l := out[i]

			if l.value >= 10 {
				left := leaf{
					depth: l.depth + 1,
					value: int(math.Floor(float64(l.value) / 2)),
					jump:  1,
				}

				right := leaf{
					depth: l.depth + 1,
					right: true,
					value: int(math.Ceil(float64(l.value) / 2)),
					jump:  1,
				}

				if !l.right {
					left.jump = l.jump + 1
				} else {
					right.jump = l.jump + 1
				}

				out = append(out, leaf{})

				copy(out[i+2:], out[i+1:])

				out[i] = left
				out[i+1] = right

				processMore = true
				break
			}
		}
	}

	return
}

func magnitude(in []leaf) int {
	out := make([]leaf, len(in))
	for i, k := range in {
		out[i] = k
	}

	for processMore := true; processMore; {
		processMore = false

		for i := 0; i < len(out)-1; i++ {
			l := out[i]

			if !l.right {
				r := out[i+1]
				if l.depth == r.depth && r.right {
					nL := leaf{
						depth: l.depth - 1,
						value: 3*l.value + 2*r.value,
					}

					if r.jump > l.jump {
						nL.right = true
						nL.jump = r.jump - 1
					} else {
						nL.jump = l.jump - 1
					}

					copy(out[i+1:], out[i+2:])

					out[i] = nL

					out = out[:len(out)-1]

					processMore = true
					break
				}
			}
		}
	}

	return out[0].value
}

func printNumber(in []leaf) {
	for i, k := range in {
		if !k.right {
			for j := 0; j < k.jump; j++ {
				fmt.Print("[")
			}
			fmt.Print(k.value)

			fmt.Print(",")
		} else {
			fmt.Print(k.value)
			for j := 0; j < k.jump; j++ {
				fmt.Print("]")
			}

			if i < len(in)-1 {
				fmt.Print(",")
			}
		}
	}
	fmt.Println()
	fmt.Println()
}
