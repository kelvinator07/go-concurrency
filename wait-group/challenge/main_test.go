package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("two")

	wg.Wait()

	if msg != "two" {
		t.Errorf("Expected to find two but not there")
	}

}

func Test_printMessage(t *testing.T) {
	// get the standard output
	stdOut := os.Stdout

	// create our standard output pipe
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "three"
	printMessage()

	// close our pipe
	_ = w.Close()

	// read out result written to our pipe
	result, _ := io.ReadAll(r)
	output := string(result)

	// reset back
	os.Stdout = stdOut

	if !strings.Contains(output, "three") {
		t.Errorf("Expected to find three but not there")
	}
}

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

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Expected to find Hello, universe!, but not there")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("Expected to find Hello, cosmos!, but not there")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected to find Hello, world!, but not there")
	}
}
