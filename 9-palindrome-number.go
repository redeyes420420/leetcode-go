func isPalindrome(x int) bool {
    
    if x < 0 {
        return false
    }
    
    var (
        reversed int
        lastDigit int
    )
    
    number := x
    for number > 0 {
        lastDigit = number % 10
        reversed = (reversed * 10) + lastDigit
        number = number / 10
    }
    
    return x == reversed
}
