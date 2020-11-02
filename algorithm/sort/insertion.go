package sort

func InsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		var j, cur int
		cur = nums[i]
		for j = i - 1; j >= 0 && cur < nums[j]; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = cur
	}
	return nums
}
