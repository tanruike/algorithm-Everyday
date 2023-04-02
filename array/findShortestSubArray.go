package array

import "math"

// 给定一个非空且只包含非负数的整数数组 nums，数组的 度 的定义是指数组里任一元素出现频数的最大值。
//
// 你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。
func FindShortestSubArray(nums []int) int {

	//用来返回结果
	ans := math.MaxInt
	//用来统计出现最多的次数
	maxLen := 0
	//用来记录没个数第一次出现的位置
	start := make(map[int]int)
	//在用一个map来记录没个数出现的次数
	time := make(map[int]int)

	for k, v := range nums {
		_, flag := time[v]
		//存在次数+！
		if flag {
			time[v]++
		} else {
			time[v] = 1
			start[v] = k
		}

		//
		if time[v] > maxLen {
			maxLen = time[v]

			ans = k - start[v] + 1
		} else if time[v] == maxLen {
			ans = minNum(ans, k-start[v]+1)
		}

	}

	return ans

}

func minNum(a, b int) int {
	if a > b {
		return b
	}
	return a
}
