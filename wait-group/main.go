package main

import (
	"fmt"
	"sync"
)

func printWord(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
	}

	wg.Add(len(words))

	for i, w := range words {
		go printWord(fmt.Sprintf("%d : %s", i, w), &wg)
	}

	wg.Wait()

	fmt.Println("The end")

}
