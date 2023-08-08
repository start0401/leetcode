package main

import (
	"log"
	"math"
)

func main() {
	log.Printf("result %v", longestPalindrome("bccb"))
}

func longestPalindrome(s string) string {
	result := s[:1]
	if len(s) == 1 {
		return s
	}
	long := make([]string, 2*len(s)+1)
	for i := 0; i < len(s); i++ {
		long[i*2] = "#"
		long[i*2+1] = string(s[i])
	}
	lps := make([]int, 2*len(s)+1)
	center := 0
	centerRight := 0
	maxLPSindex := 0
	maxLPS := 0

	for i := 0; i < len(long); i++ {
		mirrorIdx := 2*center - i
		if i < centerRight {
			lps[i] = int(math.Min(float64(lps[mirrorIdx]), float64(centerRight-i)))
		}
		for i+lps[i]+1 < len(long) && i-lps[i]-1 >= 0 && long[i+lps[i]+1] == long[i-lps[i]-1] {
			lps[i]++
		}

		if i+lps[i] > centerRight {
			center = i
			centerRight = i + lps[i]
		}
		if lps[i] >= maxLPS {
			maxLPS = lps[i]
			maxLPSindex = i
		}
	}

	l := long[maxLPSindex-int(maxLPS) : maxLPSindex+int(maxLPS)+1]
	pre := ""
	for j := 0; j < len(l); j++ {
		if l[j] != "#" {
			pre += l[j]
		}
	}
	if len(pre) > 0 {
		result = ""
		result += pre
	}
	return result
}
