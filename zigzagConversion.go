package main

import (
	"log"
	"strings"
)

func main() {
	log.Printf("convert %v", convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	result := make([]string, numRows)
	index := 0
	for i := 0; i < len(s); i++ {
		n := 2*numRows - 2
		index = i % n
		if i%n < numRows {
			result[index] += string(s[i])
		}

		if i%n >= numRows {
			result[2*numRows-index-2] += string(s[i])
		}
	}
	return strings.Join(result, "")
}
