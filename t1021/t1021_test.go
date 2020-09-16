package t1021

import "testing"

func Test1(t *testing.T) {
	t.Log(removeOuterParentheses("(()())(())"))
	t.Log(removeOuterParentheses("(()())(())(()(()))"))
	t.Log(removeOuterParentheses("()()"))
	t.Log(removeOuterParentheses(""))
}
