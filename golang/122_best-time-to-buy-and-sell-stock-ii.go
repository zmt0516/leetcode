func maxProfit(prices []int) int {
	buy := prices[0]
	r := 0
	bought := true
	for i := 1; i < len(prices); i++ {
		if !bought {
			buy = prices[i]
			bought = true
		} else {
			if prices[i] < buy {
				buy = prices[i]
			}
			if prices[i] > buy && (i+1 == len(prices) || prices[i+1] <= prices[i]) {
				r += prices[i] - buy
				bought = false
			}

		}

	}

	return r

}