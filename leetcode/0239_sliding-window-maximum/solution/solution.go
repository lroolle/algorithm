package solution

// TC: O(nk) SC: O(n-k+1)
func maxSlidingWindow(nums []int, k int) []int {
	steps := len(nums) - k + 1
	ret := make([]int, steps)

	window := make([]int, k)
	for i := 0; i < k; i++ {
		window[i] = nums[i]
	}

	for i := 0; i < steps; i++ {
		ret[i] = max(window)
		window = window[1:]

		if k+i > len(nums)-1 {
			break
		}
		window = append(window, nums[k+i])

	}
	return ret
}

func max(nums []int) int {
	m := nums[0]
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}

func maxSlidingWindowDeque(nums []int, k int) []int {
	ret := []int{}
	deque := []int{}

	// cleanDeque do things below:
	//   1. Pop left index which is outside the window
	//   2. Pop right index which num is smaller than next one
	cleanDeque := func(i int) {
		if len(deque) > 0 && deque[0] == i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[i] > nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
	}

	// Init deque
	for i := 0; i < k; i++ {
		cleanDeque(i)
		deque = append(deque, i)
	}
	ret = append(ret, nums[deque[0]])

	// Steps Move window forward
	for i := k; i < len(nums); i++ {
		cleanDeque(i)
		deque = append(deque, i)
		ret = append(ret, nums[deque[0]])
	}
	return ret
}
