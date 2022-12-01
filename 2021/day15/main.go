package main

import (
	"aoc2021/util"
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type vec struct {
	y, x, diff int
}

type PriorityQueue []vec

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].diff < pq[j].diff
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(vec))
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

var lines = util.ReadLines("2021/day15/input.txt")

var data [][]int

func init() {
	for _, t := range lines {
		s := strings.Split(t, "")

		n := make([]int, len(s))

		for i, k := range s {
			n[i], _ = strconv.Atoi(k)
		}

		data = append(data, n)
	}
}

func main() {
	fmt.Println("Part 1:", traverse(data))

	dY := len(data)
	dX := len(data[0])

	data2 := make([][]int, dY*5)
	for y := range data2 {
		data2[y] = make([]int, dX*5)

		for x := range data2[y] {
			data2[y][x] = data[y%dY][x%dX] + y/dY + x/dX
			if data2[y][x] > 9 {
				data2[y][x] -= 9
			}
		}
	}

	fmt.Println("Part 2:", traverse(data2))
}

func traverse(graph [][]int) int {
	dY := len(graph)
	dX := len(graph[0])

	dists := make([][]int, dY)
	for y := range dists {
		dists[y] = make([]int, dX)
		for x := range dists[y] {
			dists[y][x] = math.MaxInt
		}
	}

	dists[0][0] = 0

	priorityQueue := make(PriorityQueue, 0)

	heap.Init(&priorityQueue)

	heap.Push(&priorityQueue, vec{0, 0, 0})

	for len(priorityQueue) > 0 {
		el := heap.Pop(&priorityQueue).(vec)

		nds := []vec{{el.y - 1, el.x, 0}, {el.y + 1, el.x, 0}, {el.y, el.x - 1, 0}, {el.y, el.x + 1, 0}}

		for _, n := range nds {
			if !(n.y >= 0 && n.y <= dY-1 &&
				n.x >= 0 && n.x <= dX-1) {
				continue
			}

			cD := dists[el.y][el.x] + graph[n.y][n.x]

			if cD < dists[n.y][n.x] {
				dists[n.y][n.x] = cD
				n.diff = cD

				heap.Push(&priorityQueue, n)
			}
		}
	}

	return dists[dY-1][dX-1]
}
