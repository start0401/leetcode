package main

import (
	"log"
)

func main() {
	a := []string{"flower", "flow", "flight"}
	//a := []string{"flower", "flower", "flower", "flower"}
	// a := []string{"dog", "racecar", "car"}
	//a := []string{"aacc", "aa", "aa", "aa", "aaca"}
	log.Printf("longestCommonPrefix %v", longestCommonPrefix(a))
}

func longestCommonPrefix(strs []string) string {
	result := strs[0]

	if len(strs) == 1 {
		return strs[0]
	}

	var compareString func(str string, count int) bool
	compareString = func(str string, count int) bool {
		if count > len(str) {
			return false
		}
		if str[:count] == result {
			return true
		}
		return false
	}

	count := len(result)
	for i := 0; i < len(strs); i++ {
		if compareString(strs[i], count) {
			result = strs[i][:count]
		} else {
			for count > 0 {
				if count > len(strs[i]) {
					count--
					continue
				}
				result = result[:count]
				if compareString(strs[i], count) {
					break
				}
				count--
			}
		}
	}
	return result[:count]
}
