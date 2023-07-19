package main

import (
	"GoLab/account"
	"GoLab/depositSystem"
	"fmt"
)

func main() {
	acc := depositSystem.DepositSystem{
		AccountMap: make(map[string]account.Account),
	}
	acc.CreateAccount("001")
	fmt.Println(acc)
	acc.Deposit("001", 500)
	fmt.Println(acc)
	acc.WithDraw("001", 100)
	fmt.Println(acc)
	acc.CreateAccount("002")
	acc.Transfer("001", "002", 100)
	fmt.Println(acc)
}

func initRoute() {

}

func initStandradRoute() {

}
