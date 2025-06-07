package week1

func NonRepeati(nums []int) int {
	m := make(map[int]int) // nums:count
	for _, n := range nums {
		m[n]++
	}

	for _, k := range nums {
		if m[k] == 1 {
			return k
		}
	}

	return -1
}
