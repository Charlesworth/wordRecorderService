package main

import "testing"

func TestCheckFlags(t *testing.T) {
	*inputPort = 1
	*statsPort = 5000
	testErr := checkFlags()
	if testErr == nil {
		t.Error("did not return error when port set to system reserved port (<1024)")
	}

	*inputPort = 5000
	*statsPort = 5000
	testErr = checkFlags()
	if testErr == nil {
		t.Error("did not return error when input and stat port are using the same port number")
	}

	*inputPort = 5001
	*statsPort = 5000
	testErr = checkFlags()
	if testErr != nil {
		t.Error("returned an error with a valid port input")
	}
}
