package besttimetobuyandsellstock

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxProfit := 0
	minBuy := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minBuy {
			minBuy = prices[i]
		} else {
			if maxProfit < (prices[i] - minBuy) {
				maxProfit = (prices[i] - minBuy)
			}
		}
	}

	return maxProfit
}
