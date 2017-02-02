package main

import "testing"

func TestCheckFlags(t *testing.T) {
	t.Log("Test: checkFlags should error when port supplied < 1024")
	*inputPort = 1
	*statsPort = 5000
	testErr := checkFlags()
	if testErr == nil {
		t.Error("Error: did not return error when port set to system reserved port (<1024)")
	}

	t.Log("Test: checkFlags should error when ports supplied are the same")
	*inputPort = 5000
	*statsPort = 5000
	testErr = checkFlags()
	if testErr == nil {
		t.Error("Error: did not return error when input and stat port are using the same port number")
	}

	t.Log("Test: checkFlags should not error when ports supplied are valid")
	*inputPort = 5001
	*statsPort = 5000
	testErr = checkFlags()
	if testErr != nil {
		t.Error("Error: returned an error with a valid port input")
	}
}
