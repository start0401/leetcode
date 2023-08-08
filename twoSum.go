package main

import "fmt"

func twoSum(nums []int, target int) []int {
	amount := 0
	var results []int
	for k, v := range nums {
		amount = amount + v
		fmt.Println("amount %s", amount)
		if amount == target {
			ok := containsNum(results, k-1)
			if !ok {
				results = append(results, k-1)
			}
			ok = containsNum(results, k)
			if !ok {
				results = append(results, k)
			}
		}
		if k != 0 {
			amount = amount - nums[k-1]
		}
	}
	return results
}

func containsNum(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
