package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloGaes(number int) {
	fmt.Println("Hello Gaes", number)
}

func RunMukbang() {
	fmt.Println("Mari kita mukbang")
}

func ResponseChannel(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hayuk Mabar Bro"
}

func TestGoroutine(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go ResponseChannel(channel)
	data := <-channel
	for i := 0; i < 3; i++ {
		go RunHelloGaes(i)
		go RunMukbang()
		fmt.Println(data)
		fmt.Println("eits kenapa tuch")
	}
	time.Sleep(1 * time.Second)
}
