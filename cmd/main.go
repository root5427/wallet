package main

import (
	"fmt"

	"github.com/root5427/wallet/pkg/wallet"
)

func main() {
	svc := &wallet.Service{}
	account, err := svc.RegisterAccount("+992900000001")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 10)
	if err != nil {
		switch err {
		case wallet.ErrAmountMustBePositive:
			fmt.Println(err)
		case wallet.ErrAccountNotFound:
			fmt.Println(err)
		}
		return
	}

	fmt.Println(account.Balance)
}
