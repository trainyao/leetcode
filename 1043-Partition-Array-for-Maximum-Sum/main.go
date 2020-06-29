package main

func main() {

}

func maxSumAfterPartitioning(A []int, K int) int {
	c := len(A)
	dp := make([]int, c+1)
	dp[0] = 0

	for i := 1; i <= c; i++ {
		currentMax := A[i-1]
		for j := 1; j <= K && i-j >= 0; j++ {
			currentMax = max(currentMax, A[i-j])
			dp[i] = max(dp[i], dp[i-j]+currentMax*j)
		}
	}

	return dp[c]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
