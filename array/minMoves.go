package array

// 453.最小操作次数使数组元素相等
//给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。
func minMoves(nums []int) int {

	min := nums[0]

	i := 1

	var result int
	for i < len(nums) {
		num := nums[i]
		if num > min {
			result += num - min
		} else {
			result += (min - num) * i
			min = num
		}
		i++
	}

	return result

}
