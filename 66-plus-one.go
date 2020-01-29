func plusOne(digits []int) []int {
    
    n := len(digits)
    
    var carry int
    add := 1
    for i := n-1; i >= 0; i-- {
        add += carry
        sum := digits[i] + add
        carry = sum / 10
        digits[i] = sum % 10
        add = 0
    }
    if carry > 0 {
        digits = append(digits, carry)
        digits[n], digits[0] = digits[0], digits[n]
    }
    
    return digits
}
