package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
	b "github.com/go-concurrency/barbershop"
)

const (
	seatingCapacity = 10
	arrivalRate     = 100
	cutDuration     = 1000 * time.Millisecond
	timeOpen        = 3 * time.Second
)

func main() {
	// seed our random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// print welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("===========================")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create the Barbershop
	shop := b.BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")

	// add Barbers
	shop.AddBarber("Frank")
	shop.AddBarber("John")
	shop.AddBarber("Bobby")

	shopClosing := make(chan bool)
	closed := make(chan bool)

	// start the barbershop as a goroutine
	go func() {
		<-time.After(timeOpen) // block until timeopen passes
		shopClosing <- true
		shop.CloseShopForDay()
		closed <- true
	}()

	// add clients
	i := 1

	go func() {
		for {
			// get a random number with average arrival rate
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.AddClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	<-closed
}
