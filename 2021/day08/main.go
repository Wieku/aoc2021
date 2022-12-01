package main

import (
	"aoc2021/util"
	"fmt"
	"sort"
	"strings"
)

var lines = util.ReadLines("day08/input.txt")

var data [][2][]string

func main() {
	for _, t := range lines {
		lr := strings.Split(t, "|")

		n := [2][]string{
			strings.Split(strings.TrimSpace(lr[0]), " "),
			strings.Split(strings.TrimSpace(lr[1]), " "),
		}

		for i := 0; i < 2; i++ {
			for j := 0; j < len(n[i]); j++ {
				n[i][j] = srt(n[i][j]) //presort values internally (dfa -> adf), crucial in part 2
			}
		}

		data = append(data, n)
	}

	p1()
	p2()
}

func p1() {
	sum := 0

	for _, d := range data {
		for _, r := range d[1] {
			l := len(r)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				sum++
			}
		}
	}

	fmt.Println("Part 1:", sum)
}

func p2() {
	cnt := func(s, e string) (count int) { // Count how many runes from e is in s
		for _, r := range e {
			if strings.ContainsRune(s, r) {
				count++
			}
		}

		return
	}

	result := 0

	for _, d := range data {
		var digits [10]string // Resolved digit signatures

		var five, six []string

		for _, v := range d[0] { // Resolve 1, 4, 7, 8
			switch len(v) {
			case 2: // 1
				digits[1] = v
			case 4: // 4
				digits[4] = v
			case 3: // 7
				digits[7] = v
			case 7: // 8
				digits[8] = v
			case 5:
				five = append(five, v)
			case 6:
				six = append(six, v)
			}
		}

		for _, v := range six { // Resolve 0, 6, 9, needed later to resolve 5 or 2
			if cnt(v, digits[1]) != 2 { // If 1 is not in candidate, then we have 6, otherwise 9 or 0
				digits[6] = v
			} else if cnt(v, digits[4]) == 4 { // 4 is in 9
				digits[9] = v
			} else { // So the last option is 0
				digits[0] = v
			}
		}

		for _, v := range five { // Resolve 2, 3, 5
			if cnt(v, digits[1]) == 2 { // If 1 is in candidate, then we have 3, otherwise 5 or 2
				digits[3] = v
			} else if cnt(digits[6], v) == 5 { // 5 is in 6
				digits[5] = v
			} else { // So the last option is 2
				digits[2] = v
			}
		}

		m := 1000
		for _, k := range d[1] {
			for o := 0; o < 10; o++ {
				if digits[o] == k {
					result += o * m
					m /= 10
					break
				}
			}
		}
	}

	fmt.Println("Part 2:", result)
}

func srt(word string) string {
	s := []rune(word)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}
