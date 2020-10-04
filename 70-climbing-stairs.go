func climbStairs(n int) int {
    if n==1 {
        return 1
    } else if n==2 {
        return 2
    } else {
        firstCount := 1
        secondCount := 2
        for i:=3;i<=n;i++ {
            tmp := secondCount
            secondCount = firstCount + tmp
            firstCount = tmp
        }
        return secondCount
    }
}