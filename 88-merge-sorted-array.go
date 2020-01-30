func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    if n == 0 {
        return
    }
    
    // cursor for '3rd' slice
    i := n + m - 1
    // cursor for 1st slice
    i1 := m - 1
    // cursor for 2nd slice
    i2 := n - 1
    
    // Go, beginning in the end, through the 1st slice
    for ; i >= 0; i-- {
        // if element of 1st is greater than of the second one
        if i1 >= 0 && nums1[i1] > nums2[i2] {
            nums1[i] = nums1[i1]
            i1--
        } else {
            nums1[i] = nums2[i2]
            i2--
        }
        
        if i2 < 0 {
            break
        }
    }
}
