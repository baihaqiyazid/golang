package gocontext

import (
	"context"
	"fmt"
	"runtime"
	// "sync"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	todo := context.TODO()

	fmt.Println(background)
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {

	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextG.Value("g"))
	fmt.Println(contextG.Value("f"))
	fmt.Println(contextD.Value("d"))
	fmt.Println(contextD.Value("b"))
	fmt.Println(contextA.Value("a"))
}

func CreateCounter(ctx context.Context) chan int{
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default :
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	// group := sync.WaitGroup{}
	fmt.Println("Total goroutines = ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	destination := CreateCounter(ctx)

	fmt.Println("After declare destination Total goroutines = ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Before cancel Total goroutines = ", runtime.NumGoroutine())
	cancel()

	fmt.Println("Before sleep Total goroutines = ", runtime.NumGoroutine())
	time.Sleep(2 * time.Second)
	
	fmt.Println("Total goroutines = ", runtime.NumGoroutine())
}

func TestContextWithTimeOut(t *testing.T) {
	// group := sync.WaitGroup{}
	fmt.Println("Total goroutines = ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5 * time.Second)
	destination := CreateCounter(ctx)
	defer cancel()

	fmt.Println("After declare destination Total goroutines = ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	time.Sleep(2 * time.Second)
	
	fmt.Println("Total goroutines = ", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	// group := sync.WaitGroup{}
	fmt.Println("Total goroutines = ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5 * time.Second))
	destination := CreateCounter(ctx)
	defer cancel()

	fmt.Println("After declare destination Total goroutines = ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	time.Sleep(2 * time.Second)
	
	fmt.Println("Total goroutines = ", runtime.NumGoroutine())
}