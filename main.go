package main

import (
	"fmt"
	"sort"
)

// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
func main() {
	// Test cases
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("[[-1 0 1] [-1 -1 2]] =", threeSum(nums))

	nums2 := []int{0, 3, -3, 1, 2, -2}
	fmt.Println("[[-3 0 3] [-2 0 2] [-3 1 2]] =", threeSum(nums2))
}

func threeSum(nums []int) [][]int {
	var result [][]int
	triplets := make(map[[3]int]bool)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					currentTriplet := []int{nums[i], nums[j], nums[k]}
					sort.Ints(currentTriplet)
					sorted := [3]int{currentTriplet[0], currentTriplet[1], currentTriplet[2]}
					if _, ok := triplets[sorted]; !ok {
						triplets[sorted] = true
						result = append(result, currentTriplet)
					}
				}
			}
		}
	}
	return result
}
