package goroutine

import (
	"fmt"
	"strconv"
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
	time.Sleep(3 * time.Second)
	for i := 0; i < 3; i++ {
		channel <- "Hayuk Mabar Bro " + strconv.Itoa(i)
	}

}

func ReceiveChannelOnlyOut(channel <-chan string) {
	for data := range channel {
		fmt.Println(data)
	}

}

func TestGoroutine(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	for i := 0; i < 3; i++ {
		go GiveChannelOnlyIn(channel)
		go ReceiveChannelOnlyOut(channel)
		go RunHelloGaes(i)
		go RunMukbang()
		fmt.Println("eits kenapa tuch")
	}
	time.Sleep(5 * time.Second)
}
