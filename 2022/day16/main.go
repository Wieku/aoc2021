package main

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

var lines = util.ReadLines("2022/day16/input.txt")

type valve struct {
	rate int
	next []int
}

func parseValves() (int, []*valve, [][]int) {
	idMap := make(map[string]int)

	var valves []*valve
	var valvesAdj []string

	for _, line := range lines {
		var vName string
		var v valve
		var vAdj string

		fmt.Sscanf(line, "Valve %s has flow rate=%d;", &vName, &v.rate)

		vI := strings.Index(line, "valves ") + 7

		if vI == 6 {
			vI = strings.Index(line, "to valve ") + 9
		}

		vAdj = line[vI:]

		idMap[vName] = len(valves)
		valves = append(valves, &v)
		valvesAdj = append(valvesAdj, vAdj)
	}

	for i, vAdj := range valvesAdj {
		for _, nV := range strings.Split(vAdj, ", ") {
			valves[i].next = append(valves[i].next, idMap[nV])
		}
	}

	m := getDistanceMatrix(valves)

	var valves1 []*valve
	var distances [][]int

	aId := idMap["AA"]
	aIdN := aId

	//Remove all stuck valves except AA for performance
	for i := 0; i < len(valves); i++ {
		if valves[i].rate > 0 || i == aId {
			var dst2 []int

			for j := 0; j < len(valves); j++ {
				if valves[j].rate > 0 || j == aId {
					dst2 = append(dst2, m[i][j])
				}
			}

			distances = append(distances, dst2)
			valves1 = append(valves1, valves[i])
		} else if i < aId {
			aIdN--
		}
	}

	return aIdN, valves1, distances
}

func main() {
	fmt.Println("Part 1:", p1())
	fmt.Println("Part 2:", p2())
}

func p1() (sum int) {
	aId, valves, m := parseValves()

	opened := uint64(1) << aId // treat "AA" as already opened

	return solve(aId, 30, opened, valves, m)
}

func p2() (sum int) {
	aId, valves, m := parseValves()

	opened := uint64(1) << aId // treat "AA" as already opened

	return solve2(aId, aId, 26, 26, opened, valves, m)
}

func solve(cNode, time int, opened uint64, valves []*valve, distances [][]int) int {
	if time <= 1 || opened == (1<<len(valves))-1 {
		return 0
	}

	a := 0

	for nId, dist := range distances[cNode] {
		if opened&(1<<nId) == 0 {
			t := time - dist - 1

			cRate := t * valves[nId].rate

			cRate += solve(nId, t, opened|(1<<nId), valves, distances)

			if cRate > a {
				a = cRate
			}
		}
	}

	return a
}

func solve2(cNode1, cNode2, time1, time2 int, opened uint64, valves []*valve, distances [][]int) int {
	if (time1 <= 1 && time2 <= 1) || opened == (1<<len(valves))-1 {
		return 0
	}

	var v1, v2 []int

	// Grab only accessible valves
	for i := 0; i < len(valves); i++ {
		if opened&(1<<i) == 0 {
			if time1-distances[cNode1][i]-1 > 0 {
				v1 = append(v1, i)
			}

			if time2-distances[cNode2][i]-1 > 0 {
				v2 = append(v2, i)
			}
		}
	}

	if len(v1) == 0 && len(v2) == 0 {
		return 0
	}

	a := 0

	if len(v1) == 1 && len(v2) == 1 && v1[0] == v2[0] {
		nId := v1[0]

		cR1 := (time1 - distances[cNode1][nId] - 1) * valves[nId].rate
		cR2 := (time2 - distances[cNode2][nId] - 1) * valves[nId].rate

		a = util.Max(a, util.Max(cR1, cR2))
	} else if len(v1) > 0 && len(v2) == 0 {
		for _, nId1 := range v1 {
			t1 := time1 - distances[cNode1][nId1] - 1
			cR := t1 * valves[nId1].rate

			cR += solve(nId1, t1, opened|(1<<nId1), valves, distances)

			a = util.Max(a, cR)
		}
	} else if len(v1) == 0 && len(v2) > 0 {
		for _, nId2 := range v2 {
			t2 := time2 - distances[cNode2][nId2] - 1
			cR := t2 * valves[nId2].rate

			cR += solve(nId2, t2, opened|(1<<nId2), valves, distances)

			a = util.Max(a, cR)
		}
	} else {
		for _, nId1 := range v1 {
			t1 := time1 - distances[cNode1][nId1] - 1

			cR1 := t1 * valves[nId1].rate

			for _, nId2 := range v2 {
				if nId1 == nId2 {
					continue
				}

				t2 := time2 - distances[cNode2][nId2] - 1

				cR2 := t2 * valves[nId2].rate

				cR3 := cR1 + cR2 + solve2(nId1, nId2, t1, t2, opened|(1<<nId1)|(1<<nId2), valves, distances)

				a = util.Max(a, cR3)
			}
		}
	}

	return a
}

// Floyd–Warshall algorithm for APSP
func getDistanceMatrix(valves []*valve) [][]int {
	m := make([][]int, len(valves))

	for i := range m {
		m2 := make([]int, len(valves))
		for j := range m2 {
			m2[j] = 1 << 30 // below max int32 for sanity
		}

		m[i] = m2
	}

	for i, v := range valves {
		m[i][i] = 0

		for _, adj := range v.next {
			m[i][adj] = 1
			m[adj][i] = 1
		}
	}

	for k := 0; k < len(valves); k++ {
		for i := 0; i < len(valves); i++ {
			for j := 0; j < len(valves); j++ {
				if m[i][j] > m[i][k]+m[k][j] {
					m[i][j] = m[i][k] + m[k][j]
				}
			}
		}
	}

	return m
}
