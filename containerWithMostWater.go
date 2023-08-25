package main

import (
	"log"
	"math"
)

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//a := []int{1, 2}
	log.Printf("number %v", maxArea(a))
}

func maxArea(height []int) int {

	i := 0
	j := len(height) - 1
	maxArea := math.Min(float64(height[i]), float64(height[j])) * float64(j-i)
	for j > i {
		if height[j] >= height[i] && j > i {
			i++
		} else {
			j--
		}
		if math.Min(float64(height[i]), float64(height[j]))*float64(j-i) >= maxArea {
			maxArea = math.Min(float64(height[i]), float64(height[j])) * float64(j-i)
		}
	}
	return int(maxArea)
}
