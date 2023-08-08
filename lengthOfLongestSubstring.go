package main

import (
	"strings"
)

func main() {

	//fmt.Println("length : %d", lengthOfLongestSubstring("ckilbkd"))
	//fmt.Println("length : %d", lengthOfLongestSubstring("ohomm"))
	//log.Printf("length %v", lengthOfLongestSubstring("ohomm"))
	//fmt.Println("length : %d", lengthOfLongestSubstring("dvdf"))
	//fmt.Println("length : %d", lengthOfLongestSubstring("bbbbb"))
	//fmt.Println("length : %d", lengthOfLongestSubstring("pwwkew"))
	//fmt.Println("length : %d", lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}

	result := string(s[:1])
	leng := 1
	count := 0
	if len(s) > 0 {
		for i := 1; i < len(s); i++ {
			result += string(s[i])
			if strings.Contains(result[:len(result)-1], result[len(result)-1:]) {
				count++
				i = count
				result = string(s[i])
			} else {
				if leng < len(result) {
					leng = len(result)
				}
			}
		}
	}
	return leng
}
