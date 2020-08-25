package leetcode

import "fmt"

type TestCase struct {
	in  interface{}
	out interface{}
}

func (tc *TestCase) NameByIndex(index int) string {
	return fmt.Sprintf("test-case: %v", index)
}

func (tc *TestCase) ErrorMessage(got interface{}) string {
	return fmt.Sprintf("input: %v, expected: %v, actual: %v", tc.in, tc.out, got)
}
