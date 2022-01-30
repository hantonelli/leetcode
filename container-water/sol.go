package main

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	rigIdx, lefIdx := 0, len(height)-1
	var maxVolume int
	if height[rigIdx] < height[lefIdx] {
		maxVolume = height[rigIdx] * (lefIdx - rigIdx)
	} else {
		maxVolume = height[lefIdx] * (lefIdx - rigIdx)
	}

	var newVolume int
	for rigIdx != lefIdx {
		if height[rigIdx] < height[lefIdx] {
			rigIdx++
		} else {
			lefIdx--
		}
		if height[rigIdx] < height[lefIdx] {
			newVolume = height[rigIdx] * (lefIdx - rigIdx)
		} else {
			newVolume = height[lefIdx] * (lefIdx - rigIdx)
		}
		if maxVolume < newVolume {
			maxVolume = newVolume
		}
	}

	return maxVolume
}
