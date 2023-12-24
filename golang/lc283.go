package main

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != 0 {
			nums[fastIndex], nums[slowIndex] = nums[slowIndex], nums[fastIndex]
			slowIndex++
		}
	}
}
