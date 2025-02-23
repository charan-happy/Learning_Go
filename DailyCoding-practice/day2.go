/*  
This problem was asked by Uber.

Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
*/

// using division
package main

import "fmt"

func productExceptSelfWithDivision(nums []int) []int {
    // Step 1: Calculate total product
    totalProduct := 1
    for _, num := range nums {
        totalProduct *= num
    }

    // Step 2: Create result array by dividing total by each number
    result := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        result[i] = totalProduct / nums[i]
    }

    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    result := productExceptSelfWithDivision(nums)
    fmt.Println("Result with division:", result) // Output: [120 60 40 30 24]

    nums2 := []int{3, 2, 1}
    result2 := productExceptSelfWithDivision(nums2)
    fmt.Println("Result with division:", result2) // Output: [2 3 6]
}

/*  
 Edge Case
If the array contains a 0, division by zero will panic. We’d need to handle this (e.g., check for zeros), but let’s assume no zeros for now to keep it simple


Complexity:
Time: 
O(n)
 – Two separate loops over the array (one for product, one for division).
Space: 
O(n)
 – For the result array (not counting the input).
*/

// without division

package main

import "fmt"

func productExceptSelfNoDivision(nums []int) []int {
    n := len(nums)
    result := make([]int, n)

    // Step 1: Fill result with prefix products
    result[0] = 1 // No numbers to the left of index 0
    for i := 1; i < n; i++ {
        result[i] = result[i-1] * nums[i-1]
    }

    // Step 2: Multiply by suffix products
    suffix := 1 // No numbers to the right initially
    for i := n-1; i >= 0; i-- {
        result[i] *= suffix
        suffix *= nums[i] // Update suffix for the next iteration
    }

    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    result := productExceptSelfNoDivision(nums)
    fmt.Println("Result without division:", result) // Output: [120 60 40 30 24]

    nums2 := []int{3, 2, 1}
    result2 := productExceptSelfNoDivision(nums2)
    fmt.Println("Result without division:", result2) // Output: [2 3 6]
}

/*
Time Complexity : 0(n) Two passes through the array (prefix and suffix), each each 
O(n), total 0(n) 
Space Complexity : 0(1)

Edge Cases
Empty Array: Both return an empty array (reasonable default).
Single Element: [5] → [1] (product of no other numbers is 1).
Zeros:
Division solution fails (division by zero).
No-division solution works: [1, 0, 2] → [0, 2, 0].

*/
