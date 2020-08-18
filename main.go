package main

import "fmt"

// 0. initialize a min and max []int{}, and count int
// loop over input slice ...
// 1. if the current index is greater or less than the last, for either min and max. Pop off the last element
// 2. always append the current input value to both min and max lists
// 3. Get the difference between the first elements in both slices.
// 4. If the first element of both min or max is equal to input[count], pop off the first element of either
// 5. if the difference greater than limit, increment the count.
// 6. return len(nums) - count

func longestSubarray(nums []int, limit int) int {
	max, min := []int{}, []int{}
	count := 0
	for i := 0; i < len(nums); i++ {
		// if the current idx is greater or less than the last..
		// pop off the last element
		for len(max) != 0 && nums[i] > max[len(max)-1] {
			max = max[:len(max)-1]
		}
		for len(min) != 0 && nums[i] < min[len(min)-1] {
			min = min[:len(min)-1]
		}

		// always append to both. Easy
		max = append(max, nums[i])
		min = append(min, nums[i])

		// get the difference between the first elements in both slices
		// is the difference greater than limit?
		if max[0]-min[0] > limit {
			// pop off the first element if the first element of min or max if min or max == nums[count]
			if max[0] == nums[count] {
				max = max[1:]
			}
			if min[0] == nums[count] {
				min = min[1:]
			}
			count++
		}
	}
	return len(nums) - count
}
