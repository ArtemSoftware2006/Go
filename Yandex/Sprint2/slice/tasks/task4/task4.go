package main

func Join(nums1, nums2 []int) []int {
	result := make([]int, 0, len(nums1)+len(nums2))
	result = append(result, nums1...)
	result = append(result, nums2...)
	return result
}
