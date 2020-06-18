import "math/bits"

func singleNumber(nums []int) int {
	var r, sum int
	for i := 0; i < bits.UintSize; i++ {
		//for i:=0;i<64;i++{
		sum = 0
		for _, v := range nums {
			sum += (v >> i) & 1
		}
		r += (sum % 3) << i
	}
	return r

}