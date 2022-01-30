package main

func rotate(a [][]int) {
	n, aTmp := len(a)-1, 0
	for j := 0; j <= n/2; j++ {
		for i := j; i < n-j; i++ {
			aTmp = a[n-j][i]
			a[n-j][i] = a[n-i][n-j]
			a[n-i][n-j] = a[j][n-i]
			a[j][n-i] = a[i][j]
			a[i][j] = aTmp
		}
	}
}

// expected: [][]int{[]int{15, 13, 2, 5}, []int{14, 3, 4, 1}, []int{12, 6, 8, 9}, []int{16, 7, 10, 11}}
// actual  : [][]int{[]int{15, 1, 9, 5}, []int{2, 3, 4, 10}, []int{13, 6, 8, 7}, []int{16, 14, 12, 11}}
