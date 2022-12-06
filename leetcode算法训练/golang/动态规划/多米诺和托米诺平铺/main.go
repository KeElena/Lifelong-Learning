func numTilings(n int) int {
    var mod int=1e9+7
    dp:=make([][4]int,n+1)
    dp[0][3]=1
    for i:=1;i<n+1;i++{
        dp[i][0]=dp[i-1][3]%mod
        dp[i][1]=(dp[i-1][0]+dp[i-1][2])%mod
        dp[i][2]=(dp[i-1][0]+dp[i-1][1])%mod
        dp[i][3]=(dp[i-1][0]+dp[i-1][1]+dp[i-1][2]+dp[i-1][3])%mod
    }
    return dp[n][3]
}