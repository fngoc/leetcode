package easy

// Merge
// Time O(N + M)
// Memory O(1)
func Merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, len(nums1)-1

	for j >= 0 && i >= 0 {
		if nums2[j] > nums1[i] {
			nums1[k] = nums2[j]
			j -= 1
			k -= 1
		} else {
			nums1[k] = nums1[i]
			i -= 1
			k -= 1
		}
	}

	for i >= 0 {
		nums1[k] = nums1[i]
		k -= 1
		i -= 1
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		k -= 1
		j -= 1
	}
}
