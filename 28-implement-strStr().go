func strStr(haystack string, needle string) int {
    
    hl := len(haystack)
    nl := len(needle)
    
    if nl == 0 {
        return 0
    }
    
    if hl == 0 {
        return -1
    }
    
    if hl == nl && haystack == needle {
        return 0
    }
    
    for i := 0; i <= hl - nl; i++ {
        if haystack[i:i+nl] == needle {
            return i
        }
    }
    
    return -1
}
