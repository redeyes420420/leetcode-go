func romanToInt(s string) int {
    
    var m map[byte]int = map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    
    var total int
    for i, _ := range s {
        if i == len(s)-1 {
            total += m[s[i]]
            break
        }
        if m[s[i]] >= m[s[i+1]] {
            total += m[s[i]]
        } else {
            total -= m[s[i]]
            continue
        }
    }
    
    return total
}
