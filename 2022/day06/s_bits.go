package main

import "math/bits"

func p1B() int {
	return getMarkerBits(4)
}

func p2B() int {
	return getMarkerBits(14)
}

func getMarkerBits(distinct int) int {
	for i := distinct - 1; i < len(rawData2); i++ {
		var cData uint

		for j := distinct - 1; j >= 0; j-- {
			cData |= rawData2[i-j]
		}

		if bits.OnesCount(cData) == distinct {
			return i + 1
		}
	}

	return -1
}
