package main

import (
	"aoc2021/util"
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
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

func init() {
	masterNode = &node{
		isDir: true,
		size:  -1,
		name:  "/",
	}

	currentNode := masterNode

	for i := 2; i < len(lines); i++ {
		line := lines[i]

		if line[0] != '$' {
			spl := strings.Split(line, " ")

			newNode := &node{
				previous: currentNode,
				isDir:    spl[0] == "dir",
				size:     -1,
				name:     spl[1],
			}

			if spl[0] != "dir" {
				newNode.size = util.Atoi(spl[0])
			}

			currentNode.leafs = append(currentNode.leafs, newNode)
		} else if strings.HasPrefix(line, "$ cd ") {
			sec := strings.TrimPrefix(line, "$ cd ")

			if sec == "/" {
				currentNode = masterNode
			} else if sec == ".." {
				currentNode = currentNode.previous
			} else {
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
	_, nds := traverseForSizes(masterNode)

	for _, cI := range nds {
		if cI <= 100000 {
			sum += cI
		}
	}

	return sum
}

func p2() int {
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
