package gogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func helloWorld()  {
	fmt.Println("Hello World")
}

func TestGoRoutines(t *testing.T)  {
	go helloWorld()
	fmt.Println("Test")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int)  {
	fmt.Printf("Display %d \n", number)
}

func TestDisplayName(t *testing.T)  {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}