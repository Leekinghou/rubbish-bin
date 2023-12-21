package main

func removeElement(nums []int, val int) int {
	slowIndex := 0

	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != val {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}

func removeElement1(nums []int, val int) int {
	left := 0
	right := len(nums) - 1
	for ; right >= 0 && nums[right] == val; right-- {
	}

	for left <= right {
		if nums[left] == val {
			nums[left] = nums[right]
			right--
		}
		left++
		for ; right >= 0 && nums[right] == val; right-- {

		}
	}
	return left
}
