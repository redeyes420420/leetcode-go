func maxSubArray(nums []int) int {
    
    lm := 0
    gm := math.MinInt64
    
    for i := 0; i < len(nums); i++ {
        lm = max(nums[i], nums[i] + lm)
        if gm < lm {
            gm = lm
        }
    }
    
    return gm
    
}

func max (a, b int) int {
    
    if a >= b {
        return a
    }
    
    return b
}
