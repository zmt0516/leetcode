func wiggleSort(nums []int) {
	i := 0
	k := 0
	for {
		if i+1 == len(nums) {
			i = 0
			//fmt.Println(k, nums)
			if k == 0 {
				break
			}
			k = 0
		}
		if nums[i] == nums[i+1] {
			k++
			j := i + 1
			for {
				j += 2
				if j >= len(nums) {
					j = j % 2
				}
				if nums[i] != nums[j] {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}

		switch i % 2 {
		case 0:
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		case 1:
			if nums[i] < nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
		i++

	}

}