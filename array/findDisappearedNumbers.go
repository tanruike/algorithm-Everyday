package array

func FindDisappearedNumbers(nums []int) []int {

	arr := make([]int, len(nums))

	for _, k := range nums {
		arr[k-1]++
	}

	var ans []int
	for i, k := range arr {
		if k == 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
