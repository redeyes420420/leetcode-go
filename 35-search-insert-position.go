func searchInsert(nums []int, target int) int {
    
    for i, el := range nums {
        if el == target {
            return i
        } else if el > target {
            return i
        }
    }
    
    return len(nums)
}
