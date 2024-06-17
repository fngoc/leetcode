package easy

// TwoSum
// Time O(N)
// Memory O(1)
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		c := target - nums[i]

		if j, ok := m[c]; ok {
			return []int{j, i}
		}

		m[nums[i]] = i
	}
	return []int{}
}
