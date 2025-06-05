package week1

func RemoveDuplicates(nums []int) []int {
	currIndex, insertIndex := 0, 0
	for currIndex < len(nums)-1 {
		// fmt.Println("---", currIndex)
		// fmt.Println(">>>", insertIndex)
		// fmt.Println(nums[currIndex])
		// fmt.Println(nums[currIndex], "!=", nums[currIndex+1])
		if nums[currIndex] != nums[currIndex+1] {
			nums[insertIndex] = nums[currIndex]
			insertIndex++
		}
		currIndex++
	}
	return nums
}
