func twoSum(numbers []int, target int) []int {
	var r []int
	j := 1
	for i, v := range numbers {
		if v+numbers[j] == target {
			r = append(r, i+1, j+1)
			return r
		} else if v+numbers[j] > target {
			for j--; j > i; j-- {
				if v+numbers[j] == target {
					r = append(r, i+1, j+1)
					return r
				}
				if v+numbers[j] < target {
					break
				}

			}
			j = i + 1

		} else {
			for j++; j < len(numbers); j++ {
				if v+numbers[j] == target {
					r = append(r, i+1, j+1)
					return r
				}
				if v+numbers[j] > target {
					break
				}

			}
			j = len(numbers) - 1

		}
	}
	return r

}