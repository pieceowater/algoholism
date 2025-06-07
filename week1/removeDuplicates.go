package week1

func RemoveDuplicates(nums []int) []int {
	c, i := 0, 1
	for ; c <= len(nums)-2; c++ {
		if nums[c] != nums[c+1] {
			nums[i] = nums[c+1]
			i++
		}
	}
	return nums[:i]
}
