package main

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

var lines = util.ReadLines("2022/day16/input.txt")

type valve struct {
	rate int
	adj  []int
}

func parse() (int, []*valve) {
	idMap := make(map[string]int)

	var valves []*valve
	var valvesAdj []string

	for _, line := range lines {
		var vName string
		var v valve
		var vAdj string

		fmt.Sscanf(line, "Valve %s has flow rate=%d; tunnels lead to valves %s", &vName, &v.rate, &vAdj)

		idMap[vName] = len(valves)
		valves = append(valves, &v)
		valvesAdj = append(valvesAdj, vAdj)
	}

	for i, vAdj := range valvesAdj {
		for _, nV := range strings.Split(vAdj, ", ") {
			valves[i].adj = append(valves[i].adj, idMap[nV])
		}
	}

	return idMap["AA"], valves
}

func main() {
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
}

func p1() int {

	return 0
}

func passValve(vNum, timeLeft int, open bool, opened uint64, valves []*valve) int {

}

func p2() int {
	return 0
}
