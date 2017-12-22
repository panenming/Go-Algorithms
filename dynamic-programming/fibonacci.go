package main

import "fmt"

var dp = [200]int64{0}

func fibMemo(n int) int64 {
	if dp[n] != 0 {
		return dp[n]
	} else {
		if n <= 2 {
			dp[n] = 1
		} else {
			dp[n] = fibMemo(n-1) + fibMemo(n-2)
		}
		return dp[n]
	}
}

func fibBotUp(n int) int64 {
	dp := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		var f int64 = 0
		if i <= 2 {
			f = 1
		} else {
			f = dp[i-1] + dp[i-2]
		}
		dp[i] = f
	}
	fmt.Println(dp)
	return dp[n]
}

func fibarray(term int) []int {
	farr := make([]int, term)
	farr[0], farr[1] = 1, 1

	for i := 2; i < term; i++ {
		farr[i] = farr[i-1] + farr[i-2]
	}
	return farr
}

// 尾递归优化
func fibTail(a, b, n int64) int64 {
	if n == 0 {
		return a
	} else {
		return fibTail(a+b, a, n-1)
	}

}
