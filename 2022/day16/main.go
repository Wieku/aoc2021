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

func solve(currentValve, currentTime int, opened uint64, valves []*valve, distances [][]int) int {
	if currentTime <= 1 || opened == (1<<len(valves))-1 {
		return 0
	}

	maxSteam := 0

	for nextValve, dist := range distances[currentValve] {
		if opened&(1<<nextValve) == 0 {
			nextTime := currentTime - dist - 1

			steam := nextTime * valves[nextValve].rate

			steam += solve(nextValve, nextTime, opened|(1<<nextValve), valves, distances)

			maxSteam = util.Max(maxSteam, steam)
		}
	}

	return maxSteam
}

func solve2(human, elephant, currentTimeH, currentTimeE int, opened uint64, valves []*valve, distances [][]int) int {
	if (currentTimeH <= 1 && currentTimeE <= 1) || opened == (1<<len(valves))-1 {
		return 0
	}

	var valvesH, valvesE []int

	// Grab only accessible and closed valves
	for i := 0; i < len(valves); i++ {
		if opened&(1<<i) == 0 {
			if currentTimeH-distances[human][i]-1 > 0 {
				valvesH = append(valvesH, i)
			}

			if currentTimeE-distances[elephant][i]-1 > 0 {
				valvesE = append(valvesE, i)
			}
		}
	}

	if len(valvesH) == 0 && len(valvesE) == 0 { // No accessible valves left
		return 0
	}

	maxSteam := 0

	if len(valvesH) == 1 && len(valvesE) == 1 && valvesH[0] == valvesE[0] { // Same final valve, first to reach wins
		nextValve := valvesH[0]

		steamH := (currentTimeH - distances[human][nextValve] - 1) * valves[nextValve].rate
		steamE := (currentTimeE - distances[elephant][nextValve] - 1) * valves[nextValve].rate

		maxSteam = util.Max(maxSteam, util.Max(steamH, steamE))
	} else if len(valvesH) > 0 && len(valvesE) == 0 { // Only human has some closed valves left, get the best path score
		for _, nextValve := range valvesH {
			nextTime := currentTimeH - distances[human][nextValve] - 1
			steam := nextTime * valves[nextValve].rate

			steam += solve(nextValve, nextTime, opened|(1<<nextValve), valves, distances)

			maxSteam = util.Max(maxSteam, steam)
		}
	} else if len(valvesH) == 0 && len(valvesE) > 0 { // Only elephant has some closed valves left, get the best path score
		for _, nextValve := range valvesE {
			nextTime := currentTimeE - distances[elephant][nextValve] - 1
			steam := nextTime * valves[nextValve].rate

			steam += solve(nextValve, nextTime, opened|(1<<nextValve), valves, distances)

			maxSteam = util.Max(maxSteam, steam)
		}
	} else { // Try all different available valve pairs
		for _, nextValveH := range valvesH {
			nextTimeH := currentTimeH - distances[human][nextValveH] - 1
			steamH := nextTimeH * valves[nextValveH].rate

			for _, nextValveE := range valvesE {
				if nextValveH == nextValveE { // skip if valve is the same
					continue
				}

				nextTimeE := currentTimeE - distances[elephant][nextValveE] - 1
				steamE := nextTimeE * valves[nextValveE].rate

				totalSteam := steamH + steamE + solve2(nextValveH, nextValveE, nextTimeH, nextTimeE, opened|(1<<nextValveH)|(1<<nextValveE), valves, distances)

				maxSteam = util.Max(maxSteam, totalSteam)
			}
		}
	}

	return maxSteam
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
