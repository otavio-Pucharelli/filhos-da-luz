package main

import "testing"

func TestRun(t *testing.T) {
	// run the run function
	_, err := run()
	if err != nil {
		t.Errorf("failed run()")
	}
}
