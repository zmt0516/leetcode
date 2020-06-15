import "sort"

func fourSum(nums []int, target int) [][]int {
	r := [][]int{}
	fh := 1
	var a = [4]int{0, 1, len(nums) - 2, len(nums) - 1}
	sort.Ints(nums)

	for a[3] > 2 {
		//println(nums[a[0]], nums[a[1]], nums[a[2]])

		//if nums[a[0]]+nums[a[1]]+nums[a[2]] == 0 && !(ri[0] == nums[a[0]] && ri[1] == nums[a[1]] && ri[2] == nums[a[2]]) {
		if nums[a[0]]+nums[a[1]]+nums[a[2]]+nums[a[3]] == target {

			//println(nums[a[0]], nums[a[1]], nums[a[2]], nums[a[3]])
			//println("Help!\n")
			//ri[0], ri[1], ri[2] = nums[a[0]], nums[a[1]], nums[a[2]]
			r = append(r, []int{nums[a[0]], nums[a[1]], nums[a[2]], nums[a[3]]})

		}
		if fh == 1 {
			for a[1]++; a[1] < a[2]; a[1]++ {
				if nums[a[1]] != nums[a[1]-1] {
					break
				}

			}

			if a[1] >= a[2] || nums[a[0]]+nums[a[1]]+nums[a[2]]+nums[a[3]] > target {
				for a[0]++; a[0] < a[2]-1; a[0]++ {
					if nums[a[0]] != nums[a[0]-1] {
						break
					}

				}
				if a[1] >= a[2] {
					a[1] = a[2] - 1
				}
				if nums[a[0]]+nums[a[1]]+nums[a[2]]+nums[a[3]] > target {
					fh = -1
				}
				//a[1] = a[0] + 1
				if a[0] >= a[2]-1 {
					for a[2]--; a[2] > 0; a[2]-- {
						if nums[a[2]] != nums[a[2]+1] {
							break
						}

					}
					if a[2] < 2 {
						for a[3]--; a[3] > 0; a[3]-- {
							if nums[a[3]] != nums[a[3]+1] {
								break
							}

						}
						a[2] = a[3] - 1

					}
					a[0] = 0
					a[1] = 1
					fh = 1
				}
				//a[1] = a[0] + 1
				//fh = -1
			}

		} else {
			for a[1]--; a[1] > a[0]; a[1]-- {
				if nums[a[1]] != nums[a[1]+1] {
					break
				}

			}

			if a[1] <= a[0] || nums[a[0]]+nums[a[1]]+nums[a[2]]+nums[a[3]] < target {
				for a[0]++; a[0] < a[2]-1; a[0]++ {
					if nums[a[0]] != nums[a[0]-1] {
						break
					}

				}
				if a[1] <= a[0] {
					a[1] = a[0] + 1
				}
				if nums[a[0]]+nums[a[1]]+nums[a[2]]+nums[a[3]] < target {
					fh = 1
				}
				//a[1] = a[0] + 1
				if a[0] >= a[2]-1 {
					for a[2]--; a[2] > 0; a[2]-- {
						if nums[a[2]] != nums[a[2]+1] {
							break
						}

					}
					if a[2] < 2 {
						for a[3]--; a[3] > 0; a[3]-- {
							if nums[a[3]] != nums[a[3]+1] {
								break
							}

						}
						a[2] = a[3] - 1

					}
					a[0] = 0
					a[1] = 1
					fh = 1
				}

			}

		}

	}
	return r

}