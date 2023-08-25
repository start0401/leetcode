package main

import (
	"log"
	"strconv"
)

func main() {
	a := 1994
	log.Printf("result %v", intToRoman_v1(a))
}

/**
 * 效能太差
 */
func intToRoman_v1(num int) string {
	romanStr := ""
	roman := map[string]int{
		"MMM":  3000,
		"MM":   2000,
		"M":    1000,
		"CM":   900,
		"DCCC": 800,
		"DCC":  700,
		"DC":   600,
		"D":    500,
		"CD":   400,
		"CCC":  300,
		"CC":   200,
		"C":    100,
		"XC":   90,
		"LXXX": 80,
		"LXX":  70,
		"LX":   60,
		"L":    50,
		"XL":   40,
		"XXX":  30,
		"XX":   20,
		"X":    10,
		"IX":   9,
		"VIII": 8,
		"VII":  7,
		"VI":   6,
		"V":    5,
		"IV":   4,
		"III":  3,
		"II":   2,
		"I":    1,
	}

	remainders := []int{
		1: 1,
		2: 10,
		3: 100,
		4: 1000,
	}

	count := 0
	for num > 0 {
		for k, v := range roman {
			if num <= 0 {
				break
			}
			count = len(strconv.Itoa(num))
			if count == len(strconv.Itoa(v)) {
				if num/remainders[count] == v/remainders[count] {
					romanStr += k
					num = num % v
				}
			}
		}
	}

	return romanStr
}
