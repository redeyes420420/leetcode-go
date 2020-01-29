func rotate(nums []int, k int)  {
    
    n := len(nums)
    if n == 1 {
        return
    }
    k = k % n
    reverse(nums, 0, n-k-1)
    reverse(nums, n-k, n-1)
    reverse(nums, 0, n-1)
}

func reverse(arr []int, i, j int) {
    
    for i < j {
        arr[i], arr[j] = arr[j], arr[i]
        i++
        j--
    }
}
