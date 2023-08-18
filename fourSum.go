package main

import (
	"log"
	"sort"
)

func main() {
	//a := []int{2, 2, 2, 2, 2}
	a := []int{-1, 0, 1, 2, -1, -4}
	log.Printf("fourSum %v ", fourSum(a, -1))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	number := 0
	var result [][]int

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			l := j + 1
			k := len(nums) - 1
			for k > l {
				if l != j+1 && nums[l] == nums[l-1] {
					l++
					continue
				}
				number = nums[i] + nums[j] + +nums[l] + nums[k]
				if number == target {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[k]})
					l++
				}

				if number-target > 0 {
					k--
				}

				if number-target < 0 {
					l++
				}
			}
		}
	}
	return result
}
