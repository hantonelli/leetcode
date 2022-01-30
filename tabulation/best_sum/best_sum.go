package main

func BestSum(target int, numbers []int) []int {
	if target < 0 {
		return nil
	}
	if target == 0 {
		return []int{}
	}
	values := make([][]int, target+1)
	values[0] = []int{}

	for i := 0; i < target+1; i++ {
		if values[i] != nil {
			for _, num := range numbers {
				if i+num < target+1 {
					newArr := append(values[i], num)
					if values[i+num] == nil || len(newArr) < len(values[i+num]) {
						values[i+num] = newArr
					}
				}
			}
		}
	}
	return values[target]
}
