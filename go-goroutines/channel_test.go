package gogoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Yazid Baihaqi"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func SendOnly(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Yazid Baihaqi"
}

func ReceiveOnly(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestSendReceive(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go SendOnly(channel)
	go ReceiveOnly(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Muhammad"
		channel <- "Yazid"
		channel <- "Baihaqi"
	}()
	
	go func ()  {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("End")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Send data Looping " + strconv.Itoa(i)
		}

		close(channel)
	}()
	

	for data := range channel {
		fmt.Println("Receive data " + data)
	}

	fmt.Println("End")
}

func TestSelectChannel(t *testing.T)  {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data:= <- channel1:
			fmt.Println("Receive data from channel 1", data)
			counter++
		
		case data:= <- channel2:
			fmt.Println("Receive data from channel 2", data)
			counter++
		
		default :
			fmt.Println("Loading")	
		}

		if counter == 2 {
			break
		}
	}
}
