package goroutine

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

func RunHelloGaes(number int) {
	fmt.Println("Hello Gaes", number)
}

func RunMukbang() {
	fmt.Println("Mari kita mukbang")
}

func GiveChannelOnlyIn(channel chan<- string) {
	time.AfterFunc(2*time.Second, func() {
		for i := 0; i < 3; i++ {
			channel <- "Hayuk Mabar Bro " + strconv.Itoa(i)
		}
	})

}

func ReceiveChannelOnlyOut(channel <-chan string) {
	for data := range channel {
		fmt.Println(data)
	}
}

func GChan1(channel1 chan<- string) {
	time.AfterFunc(2*time.Second, func() {
		channel1 <- "lezat dan bergizi"
	})

}

func GChan2(channel2 chan<- string) {
	time.AfterFunc(2*time.Second, func() {
		channel2 <- "rasanya seperti menjadi iron men"
	})

}

func TestGoroutine(t *testing.T) {
	totalCPU := runtime.NumCPU()
	totalThread := runtime.GOMAXPROCS(-1)
	totalGoroutine := runtime.NumGoroutine()
	channel := make(chan string)
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel)
	defer close(channel1)
	defer close(channel2)

	for i := 0; i < 3; i++ {

		go GiveChannelOnlyIn(channel)
		go ReceiveChannelOnlyOut(channel)

		go RunHelloGaes(i)
		go RunMukbang()
		fmt.Println("eits kenapa tuch")
	}

	go GChan1(channel1)
	go GChan2(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println(data)
			counter++
		case data := <-channel2:
			fmt.Println(data)
			counter++
			// default:
			// 	fmt.Println("Waiting for the data ...")
		}
		if counter == 2 {
			break
		}
	}

	y := 0
	var mutex sync.Mutex
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				mutex.Lock()
				y += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("y Sebanyak = ", y)
	fmt.Println("total CPU = ", totalCPU)
	fmt.Println("total Thread = ", totalThread)
	fmt.Println("total Goroutine = ", totalGoroutine)
}
