package main

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	var accumulated int
	minHeightL, minHeightR := height[0], height[len(height)-1]
	currL, currR := 0, len(height)-1
	for currL < currR {
		if minHeightL < minHeightR {
			currL++
			if height[currL] < minHeightL {
				accumulated += minHeightL - height[currL]
			} else {
				minHeightL = height[currL]
			}
		} else {
			currR--
			if height[currR] < minHeightR {
				accumulated += minHeightR - height[currR]
			} else {
				minHeightR = height[currR]
			}
		}
	}
	return accumulated
}
