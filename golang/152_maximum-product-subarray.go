func maxProduct(nums []int) int {

	p := 1
	p2 := 0
	p3 := 0
	minus := true
	var max int
	i := 0
	for i = 0; i < len(nums); i++ {
		if nums[i] != 0 {
			break
		}
	}

	if i == len(nums) {
		return 0
	}
	max = ^int(^uint(0) >> 1)

	a := i
	b := i
	c := i
	for i < len(nums) {
		if nums[i] == 0 {
			if max < 0 {
				max = 0
			}
			if p > max {
				max = p
			}
			if p < 0 {
				if b < i-1 && p/p2 > max {
					max = p / p2
				}
				if c > a && p/p3 > max {
					max = p / p3
				}
			}
			for i < len(nums) && nums[i] == 0 {
				i++
			}
			if i == len(nums) {
				return max
			}

			minus = true
			a = i
			b = i
			c = i
			p = 1
			//fmt.Println(max,i,"!")
			continue
		}
		//fmt.Println(max,i)
		p *= nums[i]
		p3 *= nums[i]
		if nums[i] < 0 {

			if minus {
				b = i
				p2 = p
				minus = false
			}
			c = i
			p3 = nums[i]

		}
		i++

	}
	//fmt.Println(max,p,a,b,c,i)
	if p > max {
		max = p
	}
	if p < 0 {
		if b < i-1 && p/p2 > max {
			max = p / p2
		}
		if c > a && p/p3 > max {
			max = p / p3
		}
	}

	return max

}