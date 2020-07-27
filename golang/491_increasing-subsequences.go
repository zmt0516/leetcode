func findSubsequences(nums []int) [][]int {
	var r [][]int
	emap := make(map[int][]int)
	inithash := ((28629151 << 5) + 28629151)

	lenn := len(nums)
	for i := range nums {
		var ri [][]int
		ri = append(ri, []int{(inithash*199 - nums[i]), nums[i]})
		for j := i + 1; j < lenn; j++ {
			for k := len(ri) - 1; k >= 0; k-- {
				if ri[k][len(ri[k])-1] <= nums[j] {
					var rk []int
					rk = append(rk, ri[k][0]*199+nums[j])
					rk = append(rk, ri[k][1:]...)
					rk = append(rk, nums[j])
					ri = append(ri, rk)
				}

			}

		}
		//fmt.Println(ri)
		for i := 1; i < len(ri); i++ {
			emap[ri[i][0]] = ri[i][1:]
		}

	}
	for _, v := range emap {
		r = append(r, v)
	}

	return r
}