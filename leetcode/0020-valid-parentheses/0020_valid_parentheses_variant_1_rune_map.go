// https://leetcode.com/problems/valid-parentheses/

 func isValid(s string) bool {
   var stack []rune
   pairs := map[rune]rune{
        ')':'(',
        ']':'[',
        '}':'{',
   }

   for _, char := range s {
         expectedOpen , isClosed := pairs[char]

     if ! isClosed  {
         stack = append( stack, char)
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
