package main

import (
	"testing"
)

func Test_updateMessage(t *testing.T) {
	msg = "three"

	wg.Add(2)
	go updateMessage("Good day!")
	go updateMessage("Good bye!")
	wg.Wait()

	if msg != "Good day!" {
		t.Error("incorrect value in msg")
	}
}
