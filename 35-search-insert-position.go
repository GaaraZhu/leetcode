func searchInsert(nums []int, target int) int {
    for i:=0; i<len(nums); i++ {
        if nums[i] == target {
            return i
        }
        
        if nums[i] < target {
            continue
        }
        
        if nums[i] > target {
            if i==0 {
                return 0
            } else {
                return i
            }
        }
    }
    
    return len(nums)
}