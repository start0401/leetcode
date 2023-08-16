package main

import (
	"log"
	"sort"
)

func main() {
	// a := []int{-1, 0, 1, 2, -1, -4}
	a := []int{0, 0, 0, 0}
	log.Printf("threeSum %v ", threeSum(a))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	number := 0

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for k > j {

			if j != i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			number = nums[i] + nums[j] + nums[k]
			if number == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
			}

			if number > 0 {
				k--
			}

			if number < 0 {
				j++
			}
		}
	}
	return result
}
