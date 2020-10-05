func twoSum(nums []int, target int) []int {
    resultMap := make(map[int]int)
    for i, v := range nums {
        v2 := target-v
        if i2,ok:=resultMap[v2];ok {
            return []int{i, i2}
        }
        resultMap[v] = i
    }
    
    return []int{}
}