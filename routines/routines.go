package routines

import (
	"fmt"
	"sync"
)

var Wg sync.WaitGroup

func Calulate() {
	Wg.Add(10000)
	for i := 1; i <= 10000; i++ {
		go calculateSquare(i, &Wg)
	}
}

func calculateSquare(i int, wg *sync.WaitGroup) {
	//time.Sleep(1 * time.Second)
	fmt.Println(i * i)
	wg.Done()

}

// Demo Example: to show that Go Routines are running independent of each others

func Start() {
	go Process()
	fmt.Println("In start")
}

func Process() {
	fmt.Println("In process")
}
