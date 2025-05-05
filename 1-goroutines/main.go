// MY EXERCISE CHALLENGE: not using channels or sync tools (mutex,atomic,wg, etc.)
// using only goroutines and make it, wait-free & lock-free & dataRace-free

// I also pay attention to not using any import

/*

BENCHMARK RESULTS: (NOTE: I made benchmark tests in a  mac mini which has M4 chipset and 24 ram)

goos: darwin
goarch: arm64
pkg: gorutinesExercise
cpu: Apple M4

BenchmarkDeposit              4204900               287.1 ns/op            48 B/op          2 allocs/op
BenchmarkWithdraw             4168310               287.1 ns/op            48 B/op          2 allocs/op
BenchmarkTotal                4085821               292.3 ns/op            40 B/op          2 allocs/op
BenchmarkParallelMixed        6951534               170.8 ns/op            48 B/op          2 allocs/op

*/

package main

type Transaction struct {
	balance int64
}

type Account struct {
	//additional account infos ...

	tn Transaction
}

func newTransaction(b int64) Transaction {
	return Transaction{
		balance: b,
	}
}

func newAccount() Account {
	return Account{
		tn: newTransaction(0),
	}
}

func (tn Transaction) Deposit(amount int64) (Transaction, bool) {
	return Transaction{balance: tn.balance + amount}, true
}

func (tn Transaction) Withdraw(amount int64) (Transaction, bool) {
	if amount < 0 {
		// negative withdraw → deposit(|n|)
		return Transaction{balance: tn.balance - amount}, true
	}
	if tn.balance < amount {
		return tn, false
	}

	return Transaction{balance: tn.balance - amount}, true
}

func (acc Account) Deposit(amount int64, callBack func(Account, bool)) {
	local := acc
	go func() {
		newTn, success := local.tn.Deposit(amount)

		callBack(Account{tn: newTn}, success)
	}()
}

func (acc Account) Withdraw(amount int64, callBack func(Account, bool)) {
	local := acc
	go func() {
		newTn, success := local.tn.Withdraw(amount)

		callBack(Account{tn: newTn}, success)
	}()
}

func Transfer(from, to Account, amount int64, callBack func(Account, Account, bool)) {

	from.Withdraw(amount, func(newFrom Account, withdrawSuccess bool) {

		if !withdrawSuccess {
			callBack(from, to, false)
			return
		}

		to.Deposit(amount, func(newTo Account, depositSuccess bool) {
			if !depositSuccess {
				callBack(from, to, false)
				return
			}
			callBack(newFrom, newTo, true)
			return
		})
		return
	})
}

func main() {}
