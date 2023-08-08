package main

import (
	"log"
	"math"
)

func main() {
	a := []byte("   -42")
	log.Printf("a %v ", a)
	log.Printf("myAtoi %v ", myAtoi("   -42"))
	log.Printf("myAtoi %v ", myAtoi("9223372036854775808"))
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	s0 := ""
	n := 0
	for k, ch := range []byte(s) {
		if ch == ' ' {
			if s0 == "" {
				continue
			} else {
				break
			}
		}
		if s0 == "" && (s[k] == '-' || s[k] == '+') {
			s0 = s[k:]
			if len(s0) < 1 {
				return 0
			}
			continue
		} else if s0 == "" {
			s0 = "+"
		}
		ch -= '0'
		if ch > 9 {
			break
		}
		n = n*10 + int(ch)
		if n > math.MaxInt32 {
			break
		}
	}
	if s0 != "" && s0[0] == '-' {
		n = -n
	}
	if n > math.MaxInt32 {
		return math.MaxInt32
	}

	if n < math.MinInt32 {
		return math.MinInt32
	}
	return n
}
