func productExceptSelf(nums []int) []int {
    
    n := len(nums)
    left := make([]int, n)
    right := make([]int, n)
    out := make([]int, n)
    
    left[0] = 1
    right[n-1] = 1
    
    for i := 1; i < n; i++ {
        left[i] = nums[i-1] * left[i-1]
    }
    
    for i := n-2; i >= 0; i-- {
        right[i] = nums[i+1] * right[i+1]
    }
    
    for i := 0; i < n; i++ {
        out[i] = right[i] * left[i]
    }
    
    fmt.Println(left, right)
    
    return out
}
