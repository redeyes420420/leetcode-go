func fib(N int) int {
    
    if N <= 1 {
        return N
    }
    
    arr := make([]int, N+1)
    arr[0] = 0
    arr[1] = 1
    
    for i := 2; i <= N; i++ {
        arr[i] = arr[i-1] + arr[i-2]
    }
    
    return arr[N]
}
