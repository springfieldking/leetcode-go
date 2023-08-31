package solution

// isValid
func isValid(s string) bool {
	var pairs = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if left, ok := pairs[c]; ok {
			// right char pop stack
			// check has element, and first char match left char
			ls := len(stack)
			if ls == 0 {
				return false
			} else if stack[ls-1] != left {
				return false
			} else {
				stack = stack[:ls-1]
			}
		} else {
			// left char push stack
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
