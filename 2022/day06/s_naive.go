package main

func p1() int {
	return getMarker(4)
}

func p2() int {
	return getMarker(14)
}

func getMarker(distinct int) int {
	line := lines[0]

	i := distinct - 1

mainLoop:
	for ; i < len(line); i++ {

		for j := distinct - 1; j >= 1; j-- {
			for k := 0; k < j; k++ {
				if line[i-j] == line[i-k] {
					continue mainLoop
				}
			}
		}

		break
	}

	return i + 1
}
