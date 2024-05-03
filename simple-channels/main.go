package main

import (
	"fmt"
	"strings"
)

// ping <-chan receive only, pong chan<- send only
func shoot(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {

	ping := make(chan string)
	pong := make(chan string)

	go shoot(ping, pong)

	fmt.Println("Type something and press ENTER (enter q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput

		// wait for response
		response := <-pong
		fmt.Println("Response:", response)
	}

	fmt.Println("All done, Closing channels")
	close(ping)
	close(pong)
}
