package easy

// RemoveDuplicates
// Time O(N)
// Memory O(1)
func RemoveDuplicates(nums []int) int {
	prev := nums[0]
	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[k] = nums[i]
			k += 1
		}
		prev = nums[i]
	}
	return k
}
