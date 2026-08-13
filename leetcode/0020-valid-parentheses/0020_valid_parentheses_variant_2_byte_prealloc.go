package main

// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
 
    stack := make([]byte, 0, len(s))
 
    pairs := map[byte]byte{
        ')':'(',
        ']':'[',
        '}':'{',
    }
 
   for _, char := range []byte(s) {
         expectedOpen , isClosed := pairs[char]
 
      if ! isClosed  {
         stack = append(stack, char)
      } else if len(stack) == 0 {
         return false
      } else if stack[len(stack)-1] == expectedOpen {
            stack = stack[:len(stack)-1]
        } else {
            return false
        } 
      } 
   return len(stack) == 0
}
