package week1

func MoveZeros(arr []int) []int {
	insert := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			if i != insert {
				arr[insert], arr[i] = arr[i], arr[insert]
			}
			insert++
		}
	}
	return arr
}
