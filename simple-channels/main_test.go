package main

import "testing"

func Test_shoot(t *testing.T) {
	t.Parallel() // Allow parallel test execution

	// Arrange
	ping := make(chan string)
	pong := make(chan string)

	// Act
	go shoot(ping, pong)

	// Send test messages and verify responses
	tests := []struct {
		input  string
		output string
	}{
		{"hello", "HELLO!!!"},
		{"world", "WORLD!!!"},
		{"", "!!!"}, // Empty string test
	}

	// Assert
	for _, tc := range tests {
		// Send message to ping channel
		ping <- tc.input

		// Receive response from pong channel
		response := <-pong

		if response != tc.output {
			t.Errorf("Expected response: %s but got: %s", tc.output, response)
		}
	}
}
