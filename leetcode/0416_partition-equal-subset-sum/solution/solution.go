package solution

import "sort"

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	if max > target {
		return false
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 1; j <= target; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target]
}

func canPartitionDfs(nums []int) bool {
	var sum, max int
	n := len(nums)
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	target := sum / 2
	if sum%2 != 0 {
		return false
	}
	if target < max {
		return false
	}

	var dfs func([]int, int, int, [][]bool) bool
	dfs = func(nums []int, target int, i int, memo [][]bool) bool {
		if i >= n || nums[i] > target {
			return false
		}
		if nums[i] == target {
			return true
		}
		if memo[i][target] {
			return memo[i][target]
		}
		return memo[i][target] == dfs(nums, target-nums[i], i+1, memo) || dfs(nums, target, i+1, memo)
	}
	sort.Ints(nums)
	memo := make([][]bool, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]bool, target+1)
	}
	return dfs(nums, target, 0, memo)
}
