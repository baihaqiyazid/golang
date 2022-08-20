package gogoroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)


func TestGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(2 * time.Second)
			group.Done()
		}()
	}

	cpu := runtime.NumCPU()
	fmt.Println("Total CPU :", cpu)

	thread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total thread :", thread)

	goroutines := runtime.NumGoroutine()
	fmt.Println("Total goroutines :", goroutines)

	group.Wait()
}