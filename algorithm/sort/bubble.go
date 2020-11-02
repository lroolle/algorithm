package sort

func BubbleSort(nums []int) []int {
	swapped := true
	for i := 0; i < len(nums)-1 && swapped; i++ {
		swapped = false
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
	}
	return nums
}
