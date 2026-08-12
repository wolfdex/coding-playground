// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/

func removeDuplicates(s string) string {
    
    stack := []byte(s)
    writeIndex := 0

    for i:= 0; i< len(stack); i++ {

        if writeIndex > 0 && stack[writeIndex-1] == stack[i] {
            writeIndex--
        } else {
            stack[writeIndex] = stack[i]
            writeIndex++
        }
    }

    return string(stack[:writeIndex])

}
