package week1

func Intersection(nums1 []int, nums2 []int) []int { // [1, 2, 2, 1], [2, 2] -> [2,2]
	m := make(map[int]int) // num:count
	r := []int{}
	for _, num := range nums1 {
		m[num]++
	}

	for _, num := range nums2 {
		if m[num] > 0 {
			r = append(r, num)
			m[num]--
		}
	}
	return r
}
