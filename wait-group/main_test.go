package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printWord(t *testing.T) {
	// get the standard output
	stdOut := os.Stdout

	// create our standard output pipe
	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printWord("one", &wg)

	wg.Wait()

	// close our pipe
	_ = w.Close()

	// read out result written to our pipe
	result, _ := io.ReadAll(r)
	output := string(result)

	// reset back
	os.Stdout = stdOut

	if !strings.Contains(output, "one") {
		t.Errorf("Expected to find one but not there")
	}
}
