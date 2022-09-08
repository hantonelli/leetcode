package reverse_integer

import "math"

func reverse(x int) int {
	res, y := 0, x
	if y < 0 {
		y *= -1
	}
	for 0 < y {
		if (0 < x && math.MaxInt32 < res*10+y%10) || (x < 0 && (res*10+y%10)*-1 < math.MinInt32) {
			return 0
		}
		res = res * 10
		res += y % 10
		y = int(y / 10)
	}
	if x < 0 {
		return -res
	}
	return res
}
