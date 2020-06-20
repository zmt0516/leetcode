func NumberOfThree(v []int) int {
	h := len(v) - 1
	//l :=0
	m := h >> 1
	if v[m] > v[h] {
		v[m], v[h] = v[h], v[m]
	}
	if v[0] > v[h] {
		v[0], v[h] = v[h], v[0]
	}
	if v[m] > v[0] {
		v[m], v[0] = v[0], v[m]
	}
	return v[0]

}

func QuickSort(values []int, k int) {
	if len(values) <= 1 {
		return
	}
	mid, i := NumberOfThree(values), 1
	//mid , i := values[(len(values)-1)>> 1],1
	head, tail := 0, len(values)-1
	for head < tail {

		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	if len(values[head:]) < k {
		QuickSort(values[:head], k)
	}

	QuickSort(values[head+1:], k)
}

func findKthLargest(nums []int, k int) int {

	QuickSort(nums, k)

	return nums[len(nums)-k]
}