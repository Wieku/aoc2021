package main

func p1A2() int {
	return getMarkerArray2(4)
}

func p2A2() int {
	return getMarkerArray2(14)
}

func getMarkerArray2(distinct int) int {
	lastSeen := make([]int, int('z'-'a'+1))

	lastRepeat := 0

	for i := 0; i < len(rawData); i++ {
		cData := rawData[i]

		if lastSeen[cData] > lastRepeat {
			lastRepeat = lastSeen[cData]
		} else if i-lastRepeat >= distinct {
			return i + 1
		}

		lastSeen[cData] = i
	}

	return -1
}
