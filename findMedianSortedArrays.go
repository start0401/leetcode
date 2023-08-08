package main

import (
	"log"
	"math"
	"sort"
)

func main() {
	a := []int{2, 3, 5, 7, 11, 13}
	b := []int{4, 7, 8, 15, 17, 20}
	number := findMedianSortedArrays(a, b)
	log.Printf("number %v", number)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	leng := len(nums)
	flag := float64(0)
	even := false
	if leng%2 > 0 {
		flag = (float64(leng) + 1) / 2
	} else {
		even = true
		flag = ((float64(leng) / 2) + (float64(leng) / 2) + 1) / 2
	}

	avg := float64(0)
	for i := 0; i < len(nums); i++ {
		if i+1 == int(flag) && !even {
			return float64(nums[i])
		}

		if i+1 == int(flag) && even {
			avg += float64(nums[i])
		}

		if float64(i+1) == math.Ceil(flag) && even {
			avg += float64(nums[i])
		}
	}
	return avg / 2
}
