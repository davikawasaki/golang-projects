package main

import "testing"

// test hello function
func TestHello(t *testing.T) {
	defaultEmptyMsg := "Hello Dude!"
	defaultMikeMsg := "Hello Mike!"

	// test for empty argument
	emptyResult := hello("") // should return "Hello Dude!"

	if emptyResult != defaultEmptyMsg {
		t.Errorf("hello(\"\") FAILED, expected %v, got %v", defaultEmptyMsg, emptyResult)
	} else {
		t.Logf("hello(\"\") SUCCESS, expected %v, got %v", defaultEmptyMsg, emptyResult)
	}

	// test for valid argument
	result := hello("Mike") // should return "Hello Mike!"

	if result != defaultMikeMsg {
		t.Errorf("hello(\"Mike\") FAILED, expected %v, got %v", defaultMikeMsg, result)
	} else {
		t.Logf("hello(\"Mike\") SUCCESS, expected %v, got %v", defaultMikeMsg, result)
	}
}
