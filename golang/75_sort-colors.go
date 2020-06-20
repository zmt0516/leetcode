//An ugly method.

func changeColors(nums []int, n0 int, n1 int, n2 int, num int) {
	//fmt.Println(n0,n1,n2,num,"!",nums)
	if n1 < n0 {
		n1 = n0
	}
	if n1+n2 == len(nums)-1 {
		switch num {
		case 0:
			n0++
			nums[0] = 0
		case 1:
			nums[0] = 0
			nums[n0] = 1
		case 2:
			n2++
			nums[0] = 0
			nums[n0] = 1
			nums[len(nums)-n2] = 2

		}

		//fmt.Println(n0,n1,n2,num,"!",nums)
		return
	}

	if num == 0 {

		num = nums[n0+1]
		nums[n0+1] = 0
		n0++
		//n1++
		//n1--
		changeColors(nums, n0, n1, n2, num)
		return
	}
	if num == 2 {
		num = nums[len(nums)-n2-1]
		nums[len(nums)-n2-1] = 2
		n2++
		changeColors(nums, n0, n1, n2, num)
		return
	}
	if num == 1 {

		num = nums[n1+1]
		nums[n1+1] = 1
		n1++
		changeColors(nums, n0, n1, n2, num)
		return

	}

}

func sortColors(nums []int) {
	changeColors(nums, 0, 0, 0, nums[0])

}