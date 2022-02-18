package strobogrammaticnumber2

var (
	len1   = []string{"1", "0", "8"}
	len2I  = []string{"00", "11", "69", "88", "96"}
	len2E  = []string{"11", "69", "88", "96"}
	inner  = [][2]string{{"0", "0"}, {"1", "1"}, {"6", "9"}, {"8", "8"}, {"9", "6"}}
	outher = [][2]string{{"1", "1"}, {"6", "9"}, {"8", "8"}, {"9", "6"}}
)

func findStrobogrammatic(n int) []string {
	return calc(n, true)
}

func calc(n int, isExternal bool) []string {
	if n == 1 {
		return len1
	}
	if n == 2 {
		if isExternal {
			return len2E
		}
		return len2I
	}

	var prev []string
	canRest2 := 0 < int(n/2)
	if canRest2 {
		prev = calc(n-2, false)
	} else {
		prev = calc(n-1, false)
	}

	var list *[][2]string
	if isExternal {
		list = &outher
	} else {
		list = &inner
	}

	var res []string
	if canRest2 {
		for _, p := range prev {
			for _, l := range *list {
				res = append(res, l[0]+p+l[1])
			}
		}
	}
	return res
}
