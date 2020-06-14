package calc

import "testing"

// test case for Add function
func TestMathAdd(t *testing.T) {
	_, result := Add(1, 2, 3) // expected result 6

	if result != 6 {
		t.Errorf("Add(1,2,3) FAILED, expected %v but got value %v", 6, result)
	} else {
		t.Logf("Add(1,2,3) PASSED, expected %v and got value %v", 6, result)
	}
}
