package main

import (
	"aoc2021/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var lines = util.ReadLines("day17/input.txt")

var minX, maxX, minY, maxY int

func init() {
	sp := strings.Split(strings.TrimPrefix(lines[0], "target area: "), ", ")

	xStuff := strings.Split(strings.Split(sp[0], "=")[1], "..")
	yStuff := strings.Split(strings.Split(sp[1], "=")[1], "..")

	minX, _ = strconv.Atoi(xStuff[0])
	maxX, _ = strconv.Atoi(xStuff[1])

	if maxX < minX {
		minX, maxX = maxX, minX
	}

	minY, _ = strconv.Atoi(yStuff[0])
	maxY, _ = strconv.Atoi(yStuff[1])

	if maxY < minY {
		minY, maxY = maxY, minY
	}
}

func main() {
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
}

func p1() int {
	maxEndVel := -minY - 1
	return maxEndVel * (maxEndVel + 1) / 2
}

func p2() (count int) {
	minVelY := minY
	maxVelY := -minY - 1
	maxVelX := maxX

	// Shortened quadratic equation solution to find x that x(x+1)/2 = minX,
	//   then we ceil that x to get minimal velocity that terminal horizontal position is >= minX
	minVelX := int(math.Ceil((-1 + math.Sqrt(1+8*float64(minX))) / 2))

	for velX := minVelX; velX <= maxVelX; velX++ {
		for velY := minVelY; velY <= maxVelY; velY++ {
			vX, vY := velX, velY
			var x, y int
			for {
				x += vX
				y += vY

				if x >= minX && y <= maxY {
					if x <= maxX && y >= minY {
						count++
					}

					break
				}

				if vX > 0 {
					vX--
				}
				vY--
			}
		}
	}

	return
}
