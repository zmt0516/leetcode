var lenn int

func findSubsequences(nums []int) [][]int {
	var r [][]int
	lenn = len(nums)
	find(nums, 0, []int{}, &r)
	return r
}

func find(nums []int, index int, arr []int, r *[][]int) {

	if len(arr) > 1 {
		*r = append(*r, append([]int{}, arr...))
	}

	for i := index; i < lenn; i++ {
		if find2(nums, index, i) {

			continue

		}

		n := len(arr)

		if n == 0 {
			arr = append(arr, nums[i])
			find(nums, i+1, arr, r)
			arr = arr[:n]
			continue
		}
		if nums[i] >= arr[n-1] {
			arr = append(arr, nums[i])
			find(nums, i+1, arr, r)
			arr = arr[:n]
		}

	}

}
func find2(nums []int, index int, x int) bool {
	for i := index; i < x; i++ {
		if nums[i] == nums[x] {
			return true
		}
	}
	return false

}