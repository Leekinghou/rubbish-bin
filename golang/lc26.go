package main

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	return process(nums, 1)
}

func process(nums []int, k int) int {
	idx := 0
	for _, x := range nums {
		if idx < k || nums[idx-k] != x {
			nums[idx] = x
			idx++
		}
	}
	return idx
}
