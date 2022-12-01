package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"strings"
)

var lines = util.ReadLines("2021/day19/input.txt")

type vec struct {
	x, y, z int
}

func orientate(p vec, o int) vec {
	switch o / 4 {
	case 1:
		p = rotateY(p, 1)
	case 2:
		p = rotateY(p, 2)
	case 3:
		p = rotateY(p, 3)
	case 4:
		p = rotateZ(p, 1)
	case 5:
		p = rotateZ(p, 3)
	}

	return rotateX(p, o%4)
}

func rotateX(p vec, o int) vec {
	for i := 0; i < o; i++ {
		p.z, p.y = -p.y, p.z
	}

	return p
}

func rotateY(p vec, o int) vec {
	for i := 0; i < o; i++ {
		p.x, p.z = -p.z, p.x
	}

	return p
}

func rotateZ(p vec, o int) vec {
	for i := 0; i < o; i++ {
		p.x, p.y = -p.y, p.x
	}

	return p
}

var scanResults [][]vec

func init() {
	current := make([]vec, 0)

	for _, k := range lines[1:] {
		if k == "" {
			continue
		}

		if strings.HasPrefix(k, "---") {
			scanResults = append(scanResults, current)
			current = make([]vec, 0)
			continue
		}

		sp := strings.Split(k, ",")

		x, _ := strconv.Atoi(sp[0])
		y, _ := strconv.Atoi(sp[1])
		z, _ := strconv.Atoi(sp[2])

		current = append(current, vec{x, y, z})
	}

	scanResults = append(scanResults, current)
}

func main() {
	p1R, p2R := solve()
	fmt.Println("Part 1:", p1R)
	fmt.Println("Part 2:", p2R)
}

func solve() (p1R, p2R int) {
	beacons := make(map[vec]uint8)
	for _, k := range scanResults[0] {
		beacons[k] = 1
	}

	beaconArr := convToArr(beacons)

	scanners := []vec{{0, 0, 0}}

	toSolve := make([][]vec, len(scanResults)-1)
	copy(toSolve, scanResults[1:])

	for len(toSolve) > 0 {
		for tS := 0; tS < len(toSolve); tS++ {
			sc := toSolve[tS]

			for per := 0; per < 24; per++ {
				permuted := make([]vec, len(sc))
				for i, k := range sc {
					permuted[i] = orientate(k, per)
				}

				offsets := make(map[vec]uint8)

				for _, p1 := range beaconArr {
					for _, p2 := range permuted {
						offsets[vec{p1.x - p2.x, p1.y - p2.y, p1.z - p2.z}]++
					}
				}

				found := false
				var offset vec

				for p, v := range offsets {
					if v >= 12 {
						found = true
						offset = p
						break
					}
				}

				if found {
					for _, p := range permuted {
						beacons[vec{p.x + offset.x, p.y + offset.y, p.z + offset.z}] = 1
					}

					beaconArr = convToArr(beacons)

					copy(toSolve[tS:], toSolve[tS+1:])
					toSolve = toSolve[:len(toSolve)-1]
					tS--

					scanners = append(scanners, offset)

					break
				}
			}
		}
	}

	maxDist := 0

	for _, k := range scanners {
		for _, l := range scanners {
			dist := abs(k.x-l.x) + abs(k.y-l.y) + abs(k.z-l.z)

			if dist > maxDist {
				maxDist = dist
			}
		}
	}

	return len(beacons), maxDist
}

func convToArr(a map[vec]uint8) []vec {
	arr := make([]vec, len(a))

	i := 0
	for k := range a {
		arr[i] = k
		i++
	}

	return arr
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
