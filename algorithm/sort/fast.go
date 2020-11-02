package sort

func QuickSort(nums []int) []int {
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, l, r int) []int {
	if l < r {
		index := partition(nums, l, r)
		sort(nums, l, index-1)
		sort(nums, index+1, r)
	}
	return nums
}

func partition(nums []int, l, r int) int {
	p := l
	index := p + 1

	for i := index; i <= r; i++ {
		if nums[i] < nums[p] {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[p], nums[index-1] = nums[index-1], nums[p]
	return index - 1
}
