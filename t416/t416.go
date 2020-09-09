package t416

func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if sum%2 == 1 {
		return false
	}

	n := len(nums)
	sum /= 2
	dp := make([]bool, sum+1)

	dp[0] = true

	for i := 0; i < n; i++ {
		for j := sum; j >= 0; j-- {
			if j >= nums[i] {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}
	return dp[sum]
}
