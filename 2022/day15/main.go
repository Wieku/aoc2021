package main

import (
	"aoc2021/util"
	"fmt"
)

var lines = util.ReadLines("2022/day15/input.txt")

type vec2 struct {
	x, y int
}

func mDst(v1, v2 vec2) int {
	return util.Abs(v1.x-v2.x) + util.Abs(v1.y-v2.y)
}

func main() {
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
}

func p1() int {
	mark := make(map[int]uint8)

	var beacons []vec2

	for _, line := range lines {
		var s, b vec2

		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.x, &s.y, &b.x, &b.y)

		beacons = append(beacons, b)

		md := mDst(s, b)

		dstL := util.Abs(2000000 - s.y)

		if dstL <= md {
			diff := md - dstL

			for x := s.x - diff; x <= s.x+diff; x++ {
				mark[x] = 1
			}
		}
	}

	for _, b := range beacons {
		if b.y == 2000000 {
			delete(mark, b.x)
		}
	}

	return len(mark)
}

type vec2d struct {
	vec2
	d int
}

func p2() int64 {
	var sensors []vec2d
	var beacons []vec2

	for _, line := range lines {
		var s, b vec2

		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.x, &s.y, &b.x, &b.y)

		beacons = append(beacons, b)

		sensors = append(sensors, vec2d{
			vec2: s,
			d:    mDst(s, b),
		})
	}

	var vs [4]vec2

	for i, s := range sensors {

		for t := 0; t < s.d; t++ {
			vs[0].x, vs[0].y = s.x+t, s.y-s.d-1+t
			vs[1].x, vs[1].y = s.x+s.d+1-t, s.y+t
			vs[2].x, vs[2].y = s.x-t, s.y+s.d+1-t
			vs[3].x, vs[3].y = s.x-s.d-1+t, s.y+t

			for _, v := range vs {
				if v.x >= 0 && v.y >= 0 && v.x <= 4000000 && v.y <= 4000000 {
					found := false

					for j := 0; j < len(sensors); j++ {
						if j != i {
							s2 := sensors[j]

							if mDst(s2.vec2, v) <= s2.d {
								found = true
								break
							}
						}
					}

					if !found {
						return int64(v.x)*4000000 + int64(v.y)
					}
				}
			}
		}
	}

	return 0
}
