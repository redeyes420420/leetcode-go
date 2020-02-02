func maxProduct(nums []int) int {
    
    lmx := nums[0]
    lmn := nums[0]
    
    gmx := nums[0]
    
    for i := 1; i < len(nums); i++ {
        v := nums[i]
        tmp := lmx
        
        lmx = max(v, max(v * lmx, v * lmn))
        lmn = min(v, min(v * lmn, v * tmp))
        gmx = max(lmx, gmx)
    }
    
    return gmx
}

func max(a, b int) int {
    
    if a >= b {
        return a
    }
    
    return b
}

func min(a, b int) int {
    
    if a <= b {
        return a
    }
    
    return b
}
