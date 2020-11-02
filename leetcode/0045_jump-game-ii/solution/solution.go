package solution

func jump(nums []int) int {
	var jumpSteps, furthest, end int
	for i := 0; i < len(nums)-1; i++ {
		furthest = max(furthest, nums[i]+i)
		if i == end {
			end = furthest
			jumpSteps++
		}
	}
	return jumpSteps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
