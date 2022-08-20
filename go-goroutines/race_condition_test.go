package gogoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T)  {
	
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {

				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("Counter = %d", x)
}

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("Counter = %d", x)
}

//------------------------------------------------------------------------------------------------
type BankAccount struct {
	RwMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) addBalance(amount int)  {
	account.RwMutex.Lock()
	account.Balance = account.Balance + amount
	account.RwMutex.Unlock()
}

func (account *BankAccount) getBalance() int {
	account.RwMutex.RLock()
	balance := account.Balance
	account.RwMutex.RUnlock()

	return balance
}

func TestRwMutex(t *testing.T) {

	account := BankAccount{}

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.addBalance(1)
				fmt.Printf("Balance = %d \n", account.getBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("Total Balance = %d", account.getBalance())
}