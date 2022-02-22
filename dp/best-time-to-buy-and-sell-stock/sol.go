package besttimetobuyandsellstock

import "sort"

func maxProfit(prices []int) int {
	maxProfit := 0
	buyPosByValue := map[int]int{}
	sellPosByValue := map[int]int{}
	for i := 0; i < len(prices); i++ {
		if _, ok := buyPosByValue[prices[i]]; !ok {
			buyPosByValue[prices[i]] = i
		}
		sellPosByValue[prices[i]] = i
	}

	buyValues := make([]int, 0, len(sellPosByValue))
	for k := range sellPosByValue {
		buyValues = append(buyValues, k)
	}
	sort.Ints(buyValues)

	sellValues := make([]int, 0, len(sellPosByValue))
	for k := range sellPosByValue {
		sellValues = append(sellValues, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sellValues)))

	for _, buy := range buyValues {
		for _, sell := range sellValues {
			if sell < maxProfit || sell <= buy {
				break
			}
			profit := (sell - buy)
			if buyPosByValue[buy] < sellPosByValue[sell] && maxProfit < profit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
