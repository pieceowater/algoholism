package week1

func MaxNum(nums []int) int {
	if len(nums) == 0 {
		panic("empty slice")
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
