var lenp int

func maxProfit(k int, prices []int) int {

	lenp = len(prices)
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

	emap := make([]int, lenp)
	emap2 := make([]int, lenp)

	for i := 1; i <= k; i++ {
		n := 0
		for j := 1; j < lenp; j++ {
			max := emap[j]

			for ii := 0; ii < j; ii++ {
				if emap[ii]+prices[j]-prices[ii] > max {
					n++
					max = emap[ii] + prices[j] - prices[ii]
				}
			}

			if emap2[j-1] > max {
				n++
				max = emap2[j-1]
			}
			emap2[j] = max

		}
		if n != 0 {
			copy(emap, emap2)
		} else {
			return emap[lenp-1]
		}
		//fmt.Println(emap)

	}
	return emap[lenp-1]

}