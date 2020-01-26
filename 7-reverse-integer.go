// Thank you:
// https://medium.com/@ManBearPigCode/how-to-reverse-a-number-mathematically-97c556626ec6

func reverse(x int) int {
    
    var (
        reversed int
        lastDigit int
    )
    
    for x > 0 {
        lastDigit = x % 10
        reversed = (reversed * 10) + lastDigit
        x = x / 10
    }
    
    for x < 0 {
        lastDigit = x % 10
        reversed = (reversed * 10) + lastDigit
        x = x / 10
    }
    
    if reversed > math.MaxInt32 || reversed < math.MinInt32 {
        return 0
    }
    return reversed
}
