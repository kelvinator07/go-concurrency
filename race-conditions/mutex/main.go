package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock() // exclusive access
	msg = s
	m.Unlock() // release
}

func main() {
	msg = "Hello, world!"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("first", &mutex)
	go updateMessage("second", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
