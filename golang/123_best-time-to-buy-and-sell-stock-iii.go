func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	min := prices[0]
	res := 0
	res2 := 0
	gres := 0
	minp := 0
	for i := 1; i < len(prices); i++ {
		v := prices[i]
		if min > v {
			min = v
		}
		if v-min > res {
			res = v - min
		}
		if minp == 0 {
			res2, minp = findOne(prices[i+1:])
			//fmt.Println(res2,minp)
		} else {
			minp--
		}
		if res2+res > gres {
			gres = res2 + res
		}

	}
	return gres

}

func findOne(p []int) (int, int) {
	if len(p) < 2 {
		return 0, 0
	}

	min := p[0]
	minpr := 0
	minp := 0
	res := 0

	for i := 1; i < len(p); i++ {
		v := p[i]
		if min >= v {
			min = v
			minp = i
		}
		if v-min >= res {
			res = v - min
			minpr = minp
		}

	}
	return res, minpr

}