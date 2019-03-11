// https://leetcode-cn.com/problems/valid-parentheses/description/

package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func isValid(s string) bool {
	var st stack.Stack
	m := map[string]string{"(": ")", "{": "}", "[": "]"}
	for _, v := range s {
		if m[string(v)] == "" {
			if st.Peek() == nil || m[string(st.Peek().(int32))] != string(v) {
				return false
			}
			st.Pop()
		} else {
			st.Push(v)
		}
	}
	if st.Peek() != nil {
		return false
	}
	return true
}

func main() {
	b := isValid("{()}")
	fmt.Println(b)
}
