// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/

// removeDuplicates removes all adjacent duplicates using an in-place two-pointer approach.
// Memory Optimization: Reuses []byte(s) directly without allocating a secondary stack slice,
// resulting in O(1) auxiliary space and zero dynamic re-allocations during execution.

func removeDuplicates(s string) string {
    
	// Convert the string to a byte slice to modify memory in-place.
	stack := []byte(s)
	
	// writeIndex acts as our virtual stack pointer (tracks the top of the stack).
	writeIndex := 0

	// Iterate over the slice using 'i' as the read pointer.
	for i := 0; i < len(stack); i++ {
        
		// Check if the virtual stack is non-empty AND the top element (writeIndex-1)
		// matches the current character at read pointer 'i'.
		if writeIndex > 0 && stack[writeIndex-1] == stack[i] {
            
			// Pop: Move the stack pointer back. The duplicate element will be 
			// overwritten in subsequent iterations without explicit memory deletion.
			writeIndex--
            
		} else {
            
			// Push: Write the current character at the stack pointer position
			// and increment the pointer.
			stack[writeIndex] = stack[i]
			writeIndex++
		}
	}

	// Return only the valid slice range up to the virtual stack pointer.
	return string(stack[:writeIndex])
}
