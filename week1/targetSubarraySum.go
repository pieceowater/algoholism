package week1

func TargetSubarraySum(arr []int, k int) int {
	var r int              // result count
	var currSum int        // curr sum
	m := make(map[int]int) // prefix sum : count
	m[0] = 1
	for _, n := range arr {
		currSum += n
		if count, ok := m[currSum-k]; ok {
			r += count
		}
		m[currSum]++
	}
	return r
}
