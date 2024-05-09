package barbershop

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) AddBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			// if there are no clients, the barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do, so %s takes a nap.", barber)
				isSleeping = true
			}
			client, shopOpen := <-shop.ClientsChan // blocks and wait for data

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed, so send barber home and close this goroutine
				shop.sendBarberHome(barber)
				return // closes go routine
			}

		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finished cutting %s's hair.", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) CloseShopForDay() {
	color.Cyan("Closing shop for the day.")

	close(shop.ClientsChan)
	shop.Open = false

	// wait for all barbers to be done
	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan // block until all barbers sends done
	}

	close(shop.BarbersDoneChan)

	color.Green("=======================================================================")
	color.Green("The barber shop is now closed for the day and every body has gone home.")
}

func (shop *BarberShop) AddClient(client string) {
	color.Cyan("**** %s arrives!", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room.", client)
		default:
			color.Red("The waiting is full, so %s leaves.", client)
		}
	} else {
		color.Red("The barber shop is already closed so %s leaves.", client)
	}
}
