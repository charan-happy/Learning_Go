/* problem statement : 
    This problem was recently asked by Google.

Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
*/

// solution 

// bruteforce

// time complexity : 0(n^2) and space complexity 0(n)

package main

import "fmt"

func hasPairSumBruteForce(nums []int, k int) bool {
    for i := 0; i < len(nums)-1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i]+nums[j] == k {
                return true
            }
        }
    }
    return false
}

func main() {
    nums := []int{10, 15, 3, 7}
    k := 17
    result := hasPairSumBruteForce(nums, k)
    fmt.Println("Brute Force Result:", result) // Output: true
}

/*
How It Works
Use two nested loops:
Outer loop (i) iterates from the first element to the second-to-last.
Inner loop (j) iterates from i+1 to the last element.
For each pair (nums[i], nums[j]), check if their sum equals k.
Return true as soon as a valid pair is found; otherwise, return false.
Complexity
Time Complexity: 
O(n^2)
For a list of length 
n
, the outer loop runs 
n-1
 times, and the inner loop runs up to 
n-1
 times in the worst case.
Total comparisons: 
(n-1) + (n-2) + \dots + 1
, which is 
\frac{n(n-1)}{2}
, simplifying to 
O(n^2)
.
*/

// optimized approach

package main

import "fmt"

func hasPairSumOptimized(nums []int, k int) bool {
    seen := make(map[int]bool) // Map to store numbers we've seen
    for _, num := range nums {
        complement := k - num
        if seen[complement] {
            return true
        }
        seen[num] = true
    }
    return false
}

func main() {
    nums := []int{10, 15, 3, 7}
    k := 17
    result := hasPairSumOptimized(nums, k)
    fmt.Println("Optimized Result:", result) // Output: true
}

// TC : 0(n) SC : 0(n)

/* Explanation

Create a hash map (seen) to track numbers encountered.
For each number num in the list:
Calculate its complement: complement = k - num.
If complement is already in the map, we’ve found a pair that sums to k, so return true.
Otherwise, add num to the map and continue.
If no pair is found after the loop, return false.
*/