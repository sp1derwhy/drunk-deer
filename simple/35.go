package simple

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2
		num := nums[middle]
		if num == target {
			return middle
		} else if num > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}
