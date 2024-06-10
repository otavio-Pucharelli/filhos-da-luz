package main

import "testing"

func TestRun(t *testing.T) {
	// run the run function
	err := run()
	if err != nil {
		t.Error("failed run()")
	}
}
