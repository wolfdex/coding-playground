package main

// https://leetcode.com/problems/two-sum/

// twoSum finds indices of the two numbers in nums that add up to target.
// Time Complexity: O(n) - Single pass through the slice with O(1) average lookup.
// Space Complexity: O(n) - Pre-allocated map to store seen numbers and their indices.

func twoSum(nums []int, target int) []int {
  
	// Pre-allocate map capacity to avoid dynamic bucket re-hashing overhead.
	seen := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
    
		// Calculate the complement value needed to reach the target.
		check := target - nums[i]

		// Check if the complement was already encountered in previous iterations.
		if prevIndex, exists := seen[check]; exists {
			return []int{prevIndex, i}
		}

		// Store the current number with its index as key-value pair.
		seen[nums[i]] = i
	}

	return []int{}
}

/*
LEETCODE GC EXPERIMENT NOTE:
Adding the init() block forces Go's Garbage Collector to run continuously.
While it artificially drops reported memory footprint (4.87 MB -> Beats 99.95%),
it tanks execution runtime from 0ms (Beats 100%) down to 12ms (Beats 28%) due to CPU overhead.
Do NOT use in production.

func init() {
    debug.SetMemoryLimit(0)
    runtime.GC()
}
*/
