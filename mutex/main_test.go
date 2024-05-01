package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	// get the standard output
	stdOut := os.Stdout

	// create our standard output pipe
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	// close our pipe
	_ = w.Close()

	// read out result written to our pipe
	result, _ := io.ReadAll(r)
	output := string(result)

	// reset back
	os.Stdout = stdOut

	if !strings.Contains(output, "Final balance: $34320.00") {
		t.Errorf("Expected to find Final balance: $34320.00, but not there")
	}

}
