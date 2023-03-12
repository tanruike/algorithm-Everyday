package array

import (
	"sort"
)

// 665 非递减数列
// 给你一个长度为n的整数数组nums，请你判断在 最多 改变1 个元素的情况下，该数组能否变成一个非递减数列。
// 我们是这样定义一个非递减数列的：对于数组中任意的i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。
func checkPossibility(nums []int) bool {

	for i := 1; i < len(nums); i++ {
		//当前这个数破坏了生序
		if nums[i] < nums[i-1] {
			if i == 1 { //是第二个数则改变第一个数
				nums[0] = nums[1]
			} else if nums[i] >= nums[i-2] { //如果这个数大于i-2这个数 说明让i-1变成i也不会破坏前面的生序
				nums[i-1] = nums[i]
			} else { //否则如果将i-1变成i 那么i-1就小于i-2这个数了 所以让i这个数变成一个大一点的数
				nums[i] = nums[i-1]
			}
			return sort.IntsAreSorted(nums)
		}
	}

	return true

}
