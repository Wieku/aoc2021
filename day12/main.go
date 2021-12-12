package main

import (
	"aoc2021/util"
	"fmt"
	"strings"
	"unicode"
)

var lines = util.ReadLines("day12/input.txt")

var nodes map[string][]string

func main() {
	nodes = make(map[string][]string)

	for _, t := range lines {
		s := strings.Split(t, "-")

		nodes[s[0]] = append(nodes[s[0]], s[1])
		nodes[s[1]] = append(nodes[s[1]], s[0])
	}

	fmt.Println("Part 1:", traverse(false))
	fmt.Println("Part 2:", traverse(true))
}

func traverse(p2 bool) (numFinished int) {
	type tData struct {
		node string
		id   int
	}

	stack := []tData{{"start", 0}}
	visited := []map[string]uint8{{"start": 0}}

	for len(stack) > 0 {
		el := stack[0]
		stack = stack[1:]

		visited[el.id][el.node] += 1

		if unicode.IsLower(rune(el.node[0])) && visited[el.id][el.node] >= 2 {
			visited[el.id]["0"] = 1 // mark that at least one small cave was visited twice, 0 is used because all nodes are A-z
		}

		if el.node == "end" {
			numFinished++
			continue
		}

		i := 0

		for _, n := range nodes[el.node] {
			if n == "start" || (unicode.IsLower(rune(n[0])) && visited[el.id][n] >= 1 && (!p2 || visited[el.id]["0"] == 1)) {
				continue
			}

			dT := tData{n, el.id}

			if i != 0 {
				vN := make(map[string]uint8)
				for k, v := range visited[el.id] {
					vN[k] = v
				}

				visited = append(visited, vN)

				dT.id = len(visited) - 1
			}

			stack = append(stack, dT)

			i++
		}
	}

	return
}
