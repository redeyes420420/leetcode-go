func maxProfit(prices []int) int {
    
    var profit int
    min := math.MaxInt64
    
    for _, v := range prices {
        if v < min {
            min = v
        
        } else {
            profit = max(profit, v-min)
        }
    } 
    
    return profit
}

func max(a, b int) int {
    
    if a >= b {
        return a
    }
    
    return b
}
