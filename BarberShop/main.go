package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity=10
var arrivalRate=100
var cutDuration=1000*time.Millisecond
var timeOpen=5*time.Second

func main() {

	rand.Seed(time.Now().UnixNano())

	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	clientChan:=make(chan string,seatingCapacity)
	doneChan:=make(chan bool)

	//create barber shop
 
	shop:=BarberShop{
		ShopCapacity:seatingCapacity,
		HairCutDuration:cutDuration,
		NumberOfBarbers:0,
		ClientsChan:clientChan,
		BarbersDoneChan:doneChan,
		Open:true,
	}

	color.Green("The shop is open for the Day!")

	//add the barbers
	shop.addBarber("Rick")
	shop.addBarber("Frank")
	shop.addBarber("Jim")

	//start the shop
    shopClosing:=make(chan bool)
	closed:=make(chan bool)

	go func(){
       <-time.After(timeOpen)
	   shopClosing<-true
	   shop.closeShopForDay()
	   closed<-true
	}()

	//add clients
	i:=1

	go func(){
       for{
		 
		  //get a random number with average arrival rate
		  randomMilliseconds:=rand.Int()%(2*arrivalRate)
		  select{
		  case<-shopClosing:
			return
		  case <-time.After(time.Millisecond*time.Duration(randomMilliseconds)):
			shop.addClient(fmt.Sprintf("Client #%d",i))
			i++
		  }
	   }
	}()

	//block until shop is closed 
    <-closed
}