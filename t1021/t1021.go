package t1021

import (
	"strings"
)

func removeOuterParentheses(S string) string {
	sb := strings.Builder{}
	count := 0
	for _, ch := range S {
		if ch == '(' {
			count++
			if count == 1 {
				continue
			}
		} else {
			count--
			if count == 0 {
				continue
			}
		}
		sb.WriteString(string(ch))
	}
	return sb.String()
}
