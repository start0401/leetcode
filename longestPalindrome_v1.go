package main

import (
	"log"
)

func main() {
	log.Printf("result %v", longestPalindrome_v1("aacabdkacaa"))
}

func longestPalindrome_v1(s string) string {
	result := s[:1]

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= 0; j-- {
			if j < i || len(result) > len(s[i:j]) {
				break
			}
			if string(s[i]) == string(s[j]) {
				tmp := string(s[i])
				for z := 1; z <= len(s[i:j]); z++ {
					if string(s[i+z]) == string(s[j-z]) {
						tmp += string(s[i+z])
					} else {
						tmp = ""
						break
					}
				}
				if len(result) < len(tmp) {
					result = tmp
				}
			}
		}
	}

	return result
}
