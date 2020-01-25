func lengthOfLongestSubstring(s string) int {
    
    if len(s) == 1 {
        return 1
    }
    
    var (
        best int
        cur int
    )
    hash := make(map[rune]byte)
    
    for i := 0; i < len(s); i++ {
        for _, r := range s[i:] {
            if hash[r] == 0 {
                cur++
                if best < cur {
                    best = cur
                }
                hash[r] = 1
            } else {
                if best < cur {
                    best = cur
                }
                cur = 0
                hash = make(map[rune]byte)
                break
            }
        }
    }
    
    return best
}
