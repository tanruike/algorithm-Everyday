package array

func findMaxConsecutiveOnes(nums []int) int {
	zero := -1
	lengthMax := 0

	for i, k := range nums {
		if k == 0 {
			leng := i - zero - 1
			lengthMax = maxNum(lengthMax, leng)
			zero = i
		}
	}

	if nums[len(nums)-1] == 0 {
		leng := len(nums) - 1 - zero - 1
		lengthMax = maxNum(lengthMax, leng)
	} else {
		leng := len(nums) - 1 - zero
		lengthMax = maxNum(lengthMax, leng)
	}

	return lengthMax
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}

	return b
}
