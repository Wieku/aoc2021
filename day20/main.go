package main

import (
	"aoc2021/util"
	"fmt"
)

var lines = util.ReadLines("day20/input.txt")

var (
	substitution []uint8
	image        [][]uint8
)

func init() {
	substitution = parseLine(lines[0])

	for _, k := range lines[2:] {
		image = append(image, parseLine(k))
	}
}

func parseLine(in string) []uint8 {
	arr := make([]uint8, len(in))

	for i, k := range in {
		if k == '#' {
			arr[i] = 1
		}
	}

	return arr
}

func main() {
	fmt.Println("Part 1:", enchance(2))
	fmt.Println("Part 2:", enchance(50))
}

func enchance(times int) (count int) {
	img := image

	for i := 0; i < times; i++ {
		img2 := make([][]uint8, len(img)+2)
		for y := range img2 {
			img2[y] = make([]uint8, len(img[0])+2)
		}

		for y, a := range img2 {
			for x := range a {
				num := 0
				s := 8

				for y1 := y - 2; y1 <= y; y1++ {
					for x1 := x - 2; x1 <= x; x1++ {

						var pixel uint8
						if i%2 == 1 && substitution[0] == 1 { // On odd enhancements if 0=# we have to assume that the rest of the image is fully lit
							pixel = 1
						}

						if x1 >= 0 && x1 < len(img[0]) && y1 >= 0 && y1 < len(img) {
							pixel = img[y1][x1]
						}

						num |= int(pixel) << s

						s--
					}
				}

				img2[y][x] = substitution[num]
			}
		}

		img = img2
	}

	for _, a := range img {
		for _, d := range a {
			count += int(d)
		}
	}

	return
}
