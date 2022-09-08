package odd_even_jump

func oddEvenJumps(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		if i+1 == len(arr) {
			res++
			continue
		}
		valid := false
		idx := -1
		// Odd (impar)
		for j := i + 1; j < len(arr); j += 2 {
			if arr[i] <= arr[j] && (idx == -1 || arr[idx] < arr[j]) {
				idx = j
			}
		}
		if idx != -1 {
			if 0 < oddEvenJumpsInner(arr, false, idx) {
				valid = true
			}
		}
		// Even (par)
		// idx = -1
		// for j := i + 2; j < len(arr); j += 2 {
		// 	if arr[j] <= arr[0] && (idx == -1 || arr[j] < arr[idx]) {
		// 		idx = j
		// 	}
		// }
		// if idx != -1 {
		// 	if 0 < oddEvenJumpsInner(arr, true, idx) {
		// 		valid = true
		// 	}
		// }
		if valid {
			res++
		}
	}
	return res
}

func oddEvenJumpsInner(arr []int, odd bool, start int) int {
	if start+1 == len(arr) {
		return 1
	}
	idx := -1
	if odd {
		// Odd (impar)
		for j := start + 1; j < len(arr); j += 2 {
			if arr[0] <= arr[j] && (idx == -1 || arr[idx] < arr[j]) {
				idx = j
			}
		}
		if idx != -1 {
			if 0 < oddEvenJumpsInner(arr[idx:], !odd, idx) {
				return 1
			}
		}
	} else {
		// Even (par)
		idx := -1
		for j := start + 2; j < len(arr); j += 2 {
			if arr[j] <= arr[0] && (idx == -1 || arr[j] < arr[idx]) {
				idx = j
			}
		}
	}
	if idx != -1 {
		if 0 < oddEvenJumpsInner(arr[idx:], !odd, idx) {
			return 1
		}
	}
	return 0
}
