import "sort"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {

	fh := 1

	var a = [3]int{0, 1, len(nums) - 1}
	sort.Ints(nums)
	r := nums[a[0]] + nums[a[1]] + nums[a[2]]
	minimal := Abs(r - target)

	for a[2] > 1 {
		//println(nums[a[0]], nums[a[1]], nums[a[2]])

		//if nums[a[0]]+nums[a[1]]+nums[a[2]] == 0 && !(ri[0] == nums[a[0]] && ri[1] == nums[a[1]] && ri[2] == nums[a[2]]) {

		if Abs(nums[a[0]]+nums[a[1]]+nums[a[2]]-target) < minimal {

			//println(nums[a[0]], nums[a[1]], nums[a[2]])
			//println("Help!\n")
			//ri[0], ri[1], ri[2] = nums[a[0]], nums[a[1]], nums[a[2]]
			//println(r)
			r = nums[a[0]] + nums[a[1]] + nums[a[2]]
			minimal = Abs(r - target)

		}
		if fh == 1 {
			for a[1]++; a[1] < a[2]; a[1]++ {
				if nums[a[1]] != nums[a[1]-1] {
					break
				}

			}

			if a[1] >= a[2] || nums[a[0]]+nums[a[1]]+nums[a[2]]-target > 0 {
				for a[0]++; a[0] < a[2]-1; a[0]++ {
					if nums[a[0]] != nums[a[0]-1] {
						break
					}

				}
				if a[1] >= a[2] {
					a[1] = a[2] - 1
				}
				if nums[a[0]]+nums[a[1]]+nums[a[2]]-target > 0 {
					fh = -1
				}
				//a[1] = a[0] + 1
				if a[0] >= a[2]-1 {
					for a[2]--; a[2] > 0; a[2]-- {
						if nums[a[2]] != nums[a[2]+1] {
							break
						}

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

			if a[1] <= a[0] || nums[a[0]]+nums[a[1]]+nums[a[2]]-target < 0 {
				for a[0]++; a[0] < a[2]-1; a[0]++ {
					if nums[a[0]] != nums[a[0]-1] {
						break
					}

				}
				if a[1] <= a[0] {
					a[1] = a[0] + 1
				}
				if nums[a[0]]+nums[a[1]]+nums[a[2]]-target < 0 {
					fh = 1
				}
				//a[1] = a[0] + 1
				if a[0] >= a[2]-1 {
					for a[2]--; a[2] > 0; a[2]-- {
						if nums[a[2]] != nums[a[2]+1] {
							break
						}

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