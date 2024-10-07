package main

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

func (shop *BarberShop) addBarber(barber string){
    
	shop.NumberOfBarbers++

	go func(){
        isSleeping:=false
		color.Yellow("%s goes to the waiting room to check for clients",barber)

		for{
			//no client is there barber should sleep
			if(len(shop.ClientsChan)==0){
			   color.Yellow("There is nothing to do so %s takes a nap",barber)
               isSleeping=true
			}

			client,shopOpen:=<-shop.ClientsChan
            if shopOpen{
				//if barber is sleeping wake him up
				if isSleeping{
					color.Yellow("%s wakes up %s to get a haircut",client,barber)
					isSleeping=false
				}

				//cut hair
				shop.cutHair(barber,client)

			}else{
				//shop is closed send barber home and close this goroutine
                shop.sendBarberHome(barber)
				return 
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string){
     color.Green("%s barber cuts hair of %s",barber,client)
	 time.Sleep(shop.HairCutDuration)
	 color.Green("%s barber done with the haircut of %s",barber,client)
}

func (shop *BarberShop) sendBarberHome(barber string){
     color.Cyan("%s is going home",barber)
	 shop.BarbersDoneChan<-true
}

func (shop *BarberShop) closeShopForDay(){
	color.Cyan("Closing shop for a day!")
	close(shop.ClientsChan)
	shop.Open=false

	for i:=1;i<=shop.NumberOfBarbers;i++{
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)
	color.Green("------------------------------------------------------------------")
	color.Green("Barber shop is now closed for the day and everybody has gone home!")
}

func (shop *BarberShop) addClient(client string){

	color.Green("*** %s arrives",client)

	if shop.Open{
		select{
		case shop.ClientsChan<-client:
			color.Yellow("%s takes a seat and waiting",client)
		default:
			color.Red("The waiting room is full, %s leaves",client)
		}
	}else{
       color.Red("The shop is closed %s leaves",client)
	}
}