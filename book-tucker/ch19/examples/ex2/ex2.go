package ex2

import "fmt"

type account struct {
	balance int
}

// 함수
func withdrawRunc(a *account, amount int) {
	a.balance -= amount
}

// 메소드
func (a *account) withdrawRunc(amount int) {
	a.balance -= amount
}

func Ex2() {
	fmt.Println("Ex2")

	a := &account{100}

	withdrawRunc(a, 30)
	fmt.Printf("%d\n", a.balance)

	a.withdrawRunc(30)
	fmt.Printf("%d\n", a.balance)
}
