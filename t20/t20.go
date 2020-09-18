package t20

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]int32, 0, len(s))
	for _, ch := range s {
		switch ch {
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			stack = append(stack, ch)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
