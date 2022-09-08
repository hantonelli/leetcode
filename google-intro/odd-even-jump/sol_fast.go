package odd_even_jump

func oddEvenJumpsFast(A []int) int {
	numOfGoodIndices := 0
	arrLen := len(A)
	for i, _ := range A {
		jumpNum := 1
		curIndex := i
		for curIndex+1 < arrLen {
			if jumpNum%2 == 0 { //even jump
				ind := findMax(A[curIndex+1:], A[curIndex])
				if ind == -1 {
					break
				}
				curIndex = curIndex + ind
				jumpNum++
			} else { //odd jump
				ind := findMin(A[curIndex+1:], A[curIndex])
				if ind == -1 {
					break
				}
				curIndex = curIndex + ind
				jumpNum++
			}
			if curIndex == arrLen-1 {
				numOfGoodIndices++
			}

		}
	}
	return numOfGoodIndices + 1

}

func findMin(slice []int, compareVal int) int {
	//smallest value in slice > compareVal, but with smallest index if there are multiple vals
	index := len(slice)
	smallestValue := 100000
	for i, e := range slice {
		if e >= compareVal && e < smallestValue {
			if e == smallestValue && i > index {
				continue
			}
			index = i
			smallestValue = e
		}
	}
	if index == len(slice) && smallestValue == 100000 {
		return -1
	}
	return index + 1
}

func findMax(slice []int, compareVal int) int {
	//largest value in slice < compareVal, but with smallest index if there are multiple vals
	index := len(slice)
	largestValue := 0
	for i, e := range slice {
		if e <= compareVal && e > largestValue {
			if e == largestValue && i > index {
				continue
			}
			index = i
			largestValue = e
		}
	}
	if index == len(slice) && largestValue == 0 {
		return -1
	}
	return index + 1
}
