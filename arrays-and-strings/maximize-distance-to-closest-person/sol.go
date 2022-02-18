package main

import "math"

func maxDistToClosest(seats []int) int {
	lastPerson := -1
	var distance int
	maxDistance := 0

	for i := 0; i < len(seats); i++ {
		if seats[i] != 0 {
			lastPerson = i
		} else {
			if i == len(seats)-1 || lastPerson == -1 {
				distance = i - lastPerson
			} else {
				distance = int(math.Ceil(float64(i-lastPerson) / 2))
			}
			if maxDistance < distance {
				maxDistance = distance
			}
		}
	}
	return maxDistance
}
