package main

func p1A() int {
	return getMarkerArray(4)
}

func p2A() int {
	return getMarkerArray(14)
}

func getMarkerArray(distinct int) int {
	lastSeen2 := make([]int, int('z'-'a'+1))
	lastSeen := make([]int, int('z'-'a'+1))

	i := 0

mainLoop:
	for ; i < len(rawData); i++ {
		cData := rawData[i]

		lastSeen2[cData] = lastSeen[cData]
		lastSeen[cData] = i + 1

		if i >= distinct-1 {
			for j := 0; j < distinct; j++ {
				eData := rawData[i-j]
				if lastSeen2[eData]-1 >= i-distinct+1 {
					continue mainLoop
				}
			}

			break
		}
	}

	return i + 1
}
