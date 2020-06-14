package transform

import (
	"reflect"
	"testing"
)

// In package list mode, Go caches the only successful test results to avoid repeated running of the same tests.
// Whenever Go run tests on a package, Go creates a test binary and runs it.
// You can output this binary using -c flag with go test command. This will output .test file but won’t run it.
// Additionally, rename this file using -o flag.

// test case for Square function
func TestTransformSquare(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5}
	expectedResult := []int{1, 4, 9, 16, 25}

	result := SquareSlice(testSlice)

	if reflect.DeepEqual(expectedResult, result) {
		t.Log("SquareSlice PASSED")
	} else {
		t.Errorf("SquareSlice FAILED, expected %v but got %v", expectedResult, result)
	}
}
