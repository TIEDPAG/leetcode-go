package t5

import "fmt"

func longestPalindrome(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	dp := make([]int, l*l, l*l)
	min, max, maxL := 0, 0, 1
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			fmt.Println(j, l)
			if isHasPalindrome(s, i, j, dp, l) {
				if maxL < j-i+1 {
					min, max, maxL = i, j, j-i+1
					if maxL >= l {
						goto rs
					}
				}
			}
		}
	}
rs:
	fmt.Println(dp, min, max)
	return s[min : max+1]
}

func isHasPalindrome(s string, l, r int, dp []int, rowL int) bool {
	if v := dp[rowL*l+r]; v == 1 {
		return true
	}
	if l >= r {
		dp[rowL*l+r] = 1
		return true
	}
	if s[l] == s[r] {
		rs := isHasPalindrome(s, l+1, r-1, dp, rowL)
		if rs {
			dp[rowL*l+r] = 1
		}
		return rs
	}
	return false
}
