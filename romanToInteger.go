package main

import (
	"log"
)

func main() {
	a := "III"
	log.Printf("result %v", romanToInt(a))
}

func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	count := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 {
			if roman[string(s[i])] < roman[string(s[i+1])] {
				count -= roman[string(s[i])]
			} else {
				count += roman[string(s[i])]
			}
		} else {
			count += roman[string(s[i])]
		}
	}
	return count
}
