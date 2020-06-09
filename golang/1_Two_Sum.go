func twoSum(nums []int, target int) []int {
    
    var r = []int {0,1};
    
    for nums[r[0]]+nums[r[1]] !=target {
        
        if r[0]+1==r[1]{
            r[1]++
            r[0]=0
        } else {
            r[0]++
        }
        
    }
    
    return r
    
}
