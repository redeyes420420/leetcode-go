func findMin(nums []int) int {
    
    m := math.MaxInt64
    
    for _, v := range nums {
        m = min(m, v)
    }
    
    return m
}

func min(a, b int) int {
    
    if a <= b {
        return a
    }
    
    return b
}
