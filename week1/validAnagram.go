package week1

import "strings"

func ValidAnagram(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)
	letters := make(map[byte]int)
	for i := 0; i < len(str1); i++ {
		letters[str1[i]]++
		letters[str2[i]]--
	}

	for _, count := range letters {
		if count != 0 {
			return false
		}
	}

	return true
}
