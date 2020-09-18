package t20

import "testing"

func Test1(t *testing.T) {
	t.Log(isValid("()"))
	t.Log(isValid("()[]{}"))
	t.Log(isValid("(]"))
	t.Log(isValid("([)]"))
	t.Log(isValid("{[]}"))
	t.Log(isValid("]"))
}
