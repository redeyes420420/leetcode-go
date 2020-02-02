func containsDuplicate(nums []int) bool {
    
    m := make(map[int]bool, 0)
    for _, v := range nums {
        if _, f := m[v]; f {
            return true
        } else {
            m[v] = true
        }
    }
    
    return false
}
