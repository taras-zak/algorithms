package main

// dp[0] = {}
// dp[1] = {|}
// dp[2] = {||,＝}
// dp[3] = {|||, |＝, ＝|, ⎣⎤, ∟￢} = {dp[2] + {|}, dp[1] + {＝}, dp[0] + {⎣⎤, ⎡⎦}}
// dp[4] = {dp[3] + {|}, dp[2] + {＝}, dp[1] + {⎣⎤, ∟￢}, + dp[0] + {⎣–⎦, ⎡–⎤}}
func numTilings(n int) int {
	dp := make([]int, n+4)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5
	for i := 4; i <= n; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % 1_000_000_007
	}
	return dp[n]
}

func main() {

}
