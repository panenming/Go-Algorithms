package main

import "fmt"

// 最长回文子序列
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lpsRec(word string, i, j int) int {
	if i == j {
		return 1
	}
	if i > j {
		return 0
	}
	if word[i] == word[j] {
		return 2 + lpsRec(word, i+1, j-1)
	}
	return max(lpsRec(word, i, j-1), lpsRec(word, i+1, j))
}

/**
我们建立一个二维的DP数组，其中dp[i][j]表示[i,j]区间内的字符串的最长回文子序列，
那么对于递推公式我们分析一下，如果word[i]==word[j]，那么i和j就可以增加2个回文串的长度，
我们知道中间dp[i + 1][j - 1]的值，那么其加上2就是dp[i][j]的值。如果s[i] != s[j]，
那么我们可以去掉i或j其中的一个字符，然后比较两种情况下所剩的字符串谁dp值大，就赋给dp[i][j]，
那么递推公式如下：

              /  dp[i + 1][j - 1] + 2             if (s[i] == s[j])

dp[i][j] =

              \  max(dp[i + 1][j], dp[i][j - 1])  if (s[i] != s[j])
**/
func lpsDp(word string) int {
	N := len(word)
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		dp[i][i] = 1
	}
	for i := N - 1; i > 0; i-- {
		for j := i + 1; j < N; j++ {
			if word[i] == word[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[1][N-1]
}

//我们可以对空间进行优化，只用一个一维的dp数组
func lpsDp2(word string) int {
	N := len(word)
	dp := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		length := 0
		for j := i + 1; j < N; j++ {
			t := dp[j]
			if word[i] == word[j] {
				dp[j] = length + 2
			}
			length = max(length, t)
		}
	}
	maxNum := 0
	fmt.Print("dp中所有的值：")
	for _, num := range dp {
		fmt.Print(num, " ")
		if maxNum < num {
			maxNum = num
		}
	}
	fmt.Printf("\n")
	return maxNum
}
