package array

// Rotate  给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
func Rotate(nums []int, k int) {

	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}
