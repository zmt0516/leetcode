func maxProfit(k int, prices []int) int {

	lenp := len(prices)
	if lenp < 2 || k == 0 {
		return 0
	}

	if k >= lenp/2 {
		sum := 0
		for i := 1; i < lenp; i++ {
			v := prices[i] - prices[i-1]
			if v > 0 {
				sum += v
			}
		}
		return sum
	}

	buy := make([]int, k+1)
	sell := make([]int, k+1)

	MaxUint := ^uint(0)
	MaxInt := int(MaxUint >> 1)
	MinInt := -MaxInt - 1
	for i := range buy {
		buy[i] = MinInt
	}

	for _, v := range prices {

		for i := 1; i <= k; i++ {
			sell[i] = Max(buy[i]+v, sell[i])
			buy[i] = Max(sell[i-1]-v, buy[i])

		}
	}
	return sell[k]

}
func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}