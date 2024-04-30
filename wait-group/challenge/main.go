package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var msg string

func updateMessage(s string) {
	defer wg.Done()

	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage("Hello, universe!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, world!")
	wg.Wait()
	printMessage()

}
