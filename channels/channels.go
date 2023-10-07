package channels

import (
	"fmt"
	"sync"
)

// Sell sends data to the channel
func Sell(ch chan string) {
	fmt.Println("Sent data to channel")
	ch <- "Hegabovic"
}

// Buy receive data from the channel
func Buy(ch chan string) {
	fmt.Println("Waiting for Data...")
	value := <-ch
	fmt.Println("Value Received is : ", value)
}

// buffered and unbuffered channels
var ch = make(chan int, 3) // buffered channel

func SellBuff(ch chan int, wg *sync.WaitGroup) {
	ch <- 10
	ch <- 11
	ch <- 12
	// this will cause a deadlock because data are > buffer length
	//ch <- 13
	go BuyBuff(ch, wg)
	fmt.Println("sent all data to the channel")
	wg.Done()
}
func BuyBuff(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for Data.....")
	var x = <-ch
	fmt.Println("Received data", x)
	wg.Done()
}

func SellRange(ch chan int, wg *sync.WaitGroup) {
	ch <- 10
	ch <- 11
	ch <- 12
	fmt.Println("Sent all data")
	close(ch)
	wg.Done()
}

func BuyRange(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for Data...")
	for val := range ch {
		fmt.Println("Received: ", val)
	}
	wg.Done()
}
