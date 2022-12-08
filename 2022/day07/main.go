package main

import (
	"aoc2021/util"
	"fmt"
	"golang.org/x/exp/slices"
)

var lines = util.ReadLines("2022/day07/input.txt")

type node struct {
	previous *node
	leafs    []*node

	isDir bool
	size  int
	name  string
}

var masterNode *node

func parseIntUntil(st string) (side, begin int) {
	for i := 0; i < len(st); i++ {
		if st[i] == ' ' {
			return side, i + 1
		}

		side = side*10 + int(st[i]-'0')
	}

	return
}

func parse() {
	masterNode = &node{
		isDir: true,
		size:  -1,
		name:  "/",
	}

	currentNode := masterNode

	for i := 2; i < len(lines); i++ {
		line := lines[i]

		if line[0] != '$' {
			var newNode *node

			if line[0] == 'd' {
				newNode = &node{
					previous: currentNode,
					isDir:    true,
					size:     -1,
					name:     line[4:],
				}
			} else {
				siz, nPos := parseIntUntil(line)

				newNode = &node{
					previous: currentNode,
					size:     siz,
					name:     line[nPos:],
				}
			}

			currentNode.leafs = append(currentNode.leafs, newNode)
		} else if line[2] == 'c' {
			if line[5] == '/' {
				currentNode = masterNode
			} else if line[5] == '.' {
				currentNode = currentNode.previous
			} else {
				sec := line[5:]

				for _, cn := range currentNode.leafs {
					if cn.name == sec {
						currentNode = cn
						break
					}
				}
			}
		}
	}
}

func main() {
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
}

func p1() (sum int) {
	parse()
	_, nds := traverseForSizes(masterNode)

	for _, cI := range nds {
		if cI <= 100000 {
			sum += cI
		}
	}

	return sum
}

func p2() int {
	parse()
	sum, nds := traverseForSizes(masterNode)

	over := sum - (70000000 - 30000000)

	slices.SortFunc(nds, func(i, j int) bool { return i < j })

	for _, cI := range nds {
		if cI >= over {
			return cI
		}
	}

	return sum
}

func traverseForSizes(currentNode *node) (sum int, nds []int) {
	sum = traverseForSizes2(currentNode, &nds)
	return
}

func traverseForSizes2(currentNode *node, sizes *[]int) (sum int) {
	if currentNode.isDir {
		for _, cn := range currentNode.leafs {
			if cn.isDir {
				sum += traverseForSizes2(cn, sizes)
			} else {
				sum += cn.size
			}
		}

		*sizes = append(*sizes, sum)
	}

	return
}
