func search(nums []int, target int) int {
    
    if len(nums) == 0 {
        return -1
    }
    
    if len(nums) == 1 && nums[0] != target {
        return -1
    }
    
    if len(nums) == 1 && nums[0] == target {
        return 0
    }
    
    minI := minIndex(nums)
    k := len(nums) - minI
    rotate(nums, k)

    i := binarySearch(nums, 0, len(nums)-1, target)
    if i == -1 {
        return i
    }
    return ((i + len(nums) - k) % len(nums))
}

func rotate(arr []int, k int) {
    
    n := len(arr)
    if n < 2 {
        return
    }
    k = k % n
    
    reverse(arr, 0, n-k-1)
    reverse(arr, n-k, n-1)
    reverse(arr, 0, n-1)
}

func reverse(arr []int, i, j int) {
    
    for i < j {
        arr[i], arr[j] = arr[j], arr[i]
        i++
        j--
    }
    
}

func binarySearch(arr []int, low, high, target int) int {
    
    for low <= high {
        mid := (low + high) / 2
        if arr[mid] == target {
            return mid
            
        } else if arr[mid] < target {
            low = mid+1
            
        } else {
            high = mid-1
        }
    }
    
    return -1
}

func minIndex(arr []int) int {
    
    min := math.MaxInt64
    minI := -1
    for i, r := range arr {
        if r < min {
            min = r
            minI = i
        }
    }
    
    return minI
}
