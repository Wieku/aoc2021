package main

import (
	"aoc2021/util"
	"fmt"
	"strconv"
)

var lines = util.ReadLines("day16/input.txt")

type bitfield []uint64

func (b bitfield) at(i int) int {
	return int((b[i/64] >> (63 - i%64)) & 1)
}

var bits bitfield

func init() {
	s := 16
	for i := 0; i < len(lines[0]); i += s {
		r := len(lines[0]) - i
		if r > s {
			r = s
		}

		o, _ := strconv.ParseUint(lines[0][i:i+r], 16, 64)

		o <<= (s - r) * 4

		bits = append(bits, o)
	}
}

func main() {
	l, _, v := parsePacket(0)
	fmt.Println("Part 1:", l)
	fmt.Println("Part 2:", v)
}

func parsePacket(i int) (versions, offset, value int) {
	versions = bits.at(i)<<2 | bits.at(i+1)<<1 | bits.at(i+2)
	typ := bits.at(i+3)<<2 | bits.at(i+4)<<1 | bits.at(i+5)

	offset = i + 6

	switch typ {
	case 4:
		for {
			value <<= 4
			value |= bits.at(offset+1)<<3 | bits.at(offset+2)<<2 | bits.at(offset+3)<<1 | bits.at(offset+4)

			offset += 5

			if bits.at(offset-5) == 0 {
				break
			}
		}
	default:
		sLen := 15
		if bits.at(offset) == 1 {
			sLen = 11
		}

		offset++

		ln := 0
		for j := 0; j < sLen; j++ {
			ln |= bits.at(offset+j) << (sLen - 1 - j)
		}

		offset += sLen

		var packValues []int

		progress := func(sI int) (sO int) {
			var sVer, sVal int
			sVer, sO, sVal = parsePacket(sI)

			packValues = append(packValues, sVal)

			versions += sVer

			return
		}

		if sLen == 11 {
			for j := 0; j < ln; j++ {
				offset = progress(offset)
			}
		} else {
			sOff := offset
			offset += ln
			for sOff != offset {
				sOff = progress(sOff)
			}
		}

		for j, v := range packValues {
			if j == 0 {
				value = v
				continue
			}

			switch typ {
			case 0:
				value += v
			case 1:
				value *= v
			case 2, 3:
				if typ == 2 && v < value ||
					typ == 3 && v > value {
					value = v
				}
			case 5, 6, 7:
				if typ == 5 && v < value ||
					typ == 6 && v > value ||
					typ == 7 && v == value {
					value = 1
				} else {
					value = 0
				}
			}
		}
	}

	return
}
