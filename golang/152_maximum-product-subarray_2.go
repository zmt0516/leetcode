//There is no difference between this one and the first one.
//However, this one is more ingenious.

func maxProduct(nums []int) int {
	lmax := nums[0]
	lmin := nums[0]
	gmax := nums[0]
	for i := 1; i < len(nums); i++ {
		lmax, lmin = max(nums[i], nums[i]*lmin, nums[i]*lmax), min(nums[i], nums[i]*lmin, nums[i]*lmax)
		if lmax > gmax {
			gmax = lmax
		}

	}
	return gmax

}

func max(a int, b int, c int) int {
	if b > a {
		a = b
	}
	if c > a {
		a = c
	}
	return a

}

func min(a int, b int, c int) int {
	if b < a {
		a = b
	}
	if c < a {
		a = c
	}
	return a

}