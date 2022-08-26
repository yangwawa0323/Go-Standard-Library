package sync_mutex

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

var (
	mutex   sync.Mutex
	wg      sync.WaitGroup
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Fprintf(os.Stdout, "Depositing %d to account with balance %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Fprintf(os.Stdout, "Withdrawing %d from account with balance %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}

func Test_Deposit_Withdraw(t *testing.T) {
	t.Log("deposit/withdraw demo")

	balance = 1000

	wg.Add(2)
	go deposit(500, &wg)
	go withdraw(700, &wg)
	wg.Wait()

	t.Logf("New balance %d", balance)
}
