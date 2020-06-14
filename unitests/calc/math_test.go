package calc

import "testing"

// define input-result struct type
type TestDataItem struct {
	inputs   []int // inputs to `Add` function
	result   int   // result of `Add` function
	hasError bool  // does `Add` function returns error
}

// test case for Add function
func TestMathAdd(t *testing.T) {
	// _, result := Add(1, 2, 3) // expected result 6

	// if result != 6 {
	// 	t.Errorf("Add(1,2,3) FAILED, expected %v but got value %v", 6, result)
	// } else {
	// 	t.Logf("Add(1,2,3) PASSED, expected %v and got value %v", 6, result)
	// }

	// input-result data items
	dataItems := []TestDataItem{
		{[]int{1, 2, 3}, 6, false},
		{[]int{99, 99}, 198, false},
		{[]int{1, 1, 5, 5, 6}, 18, false},
		{[]int{1}, 0, true},
	}

	for _, item := range dataItems {
		err, result := Add(item.inputs...) // get result of `Add` function

		if item.hasError {
			// expected an error
			if err == nil {
				t.Errorf("Add() with args %v FAILED, expected an error but got value '%v'", item.inputs, result)
			} else {
				t.Logf("Add() with args %v PASSED, expected an error and got an error '%v'", item.inputs, err)
			}
		} else {
			// expected a value
			if result != item.result {
				t.Errorf("Add() with args %v FAILED, expected %v but got value '%v'", item.inputs, item.result, result)
			} else {
				t.Logf("Add() with args %v PASSED, expected %v and got value '%v'", item.inputs, item.result, err)
			}
		}
	}
}
