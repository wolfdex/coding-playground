// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
 
    pruefStapel := make([]byte, 0, len(s))
 
    klammernpaare := map[byte]byte{
        ')':'(',
        ']':'[',
        '}':'{',
    }
 
   for _, char := range []byte(s) {
         expectedOpen , isClosed := klammernpaare[char]
 
      if ! isClosed  {
         pruefStapel = append( pruefStapel, char)
      } else if len(pruefStapel) == 0 {
         return false
      } else if pruefStapel[len(pruefStapel)-1] == expectedOpen {
            pruefStapel = pruefStapel[:len(pruefStapel)-1]
        } else {
            return false
        } 
      } 
   return len(pruefStapel) == 0
}
