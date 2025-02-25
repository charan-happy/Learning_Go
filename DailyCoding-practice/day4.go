/*
This problem was asked by Stripe.

Given an array of integers, find the first missing positive integer in linear time and constant space. In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

You can modify the input array in-place
*/

package main

import "fmt"

func firstMissingPositive(nums []int) int {
	n := len(nums)

	// Step 1: Place each number in its correct position
	for i := 0; i < n; i++ {
		// While current number is in range [1, n] and not in its correct spot
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			// Swap it to its correct position (nums[i]-1)
			correctPos := nums[i] - 1
			nums[i], nums[correctPos] = nums[correctPos], nums[i]
		}
	}

	// Step 2: Find the first missing positive
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// If all numbers 1 to n are present, return n+1
	return n + 1
}

func main() {
	// Test cases
	tests := [][]int{
		{3, 4, -1, 1}, // Expected: 2
		{1, 2, 0},     // Expected: 3
		{7, 8, 9, 11}, // Expected: 1
		{1, 1, 1},     // Expected: 2
		{},            // Expected: 1
	}

	for _, test := range tests {
		result := firstMissingPositive(test)
		fmt.Printf("Input: %v, First missing positive: %d\n", test, result)
	}
}

// Output

// Time complexity: O(n)
// Space complexity: O(1)
// Ref: https://x.com/i/grok?conversation=1876532656982372670
