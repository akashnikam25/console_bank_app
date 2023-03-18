package main

import "fmt"

var count int = 1000

type Account struct {
	Name    string
	Number  int
	Type    string
	Balance float64
}
type Bank struct {
	//index    int
	Accounts []Account
}

func NewBank() *Bank {
	return &Bank{
		Accounts: []Account{},
	}
}

func NewAccount() Account {
	var name string
	var accountType string
	var balance float64

	fmt.Print("Enter account name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter account type: ")
	fmt.Scanln(&accountType)

	fmt.Print("Enter account balance: ")
	fmt.Scanln(&balance)

	return Account{
		Name:    name,
		Type:    accountType,
		Balance: balance,
	}
}

func (bank *Bank) createAccount(acc *Account) {
	count++
	acc.Number = count
	bank.Accounts = append(bank.Accounts, *acc)
}

func (bank *Bank) deposit(accNo int, amount float64) {
	for index, account := range bank.Accounts {
		if account.Number == accNo {
			bank.Accounts[index].Balance += amount
			break
		}
	}
}
func (bank *Bank) getBalance(accNo int) float64 {

	for _, account := range bank.Accounts {
		if account.Number == accNo {
			return account.Balance
		}

	}
	return 0.0
}
func (bank *Bank) withdraw(accNo int, amount float64) {
	for index, account := range bank.Accounts {
		if account.Number == accNo {
			bank.Accounts[index].Balance -= amount
			break
		}
	}
}
func (bank *Bank) printAccountDetails(accNo int) (Account, error) {
	for _, account := range bank.Accounts {
		if account.Number == accNo {
			return account, nil
		}
	}
	return Account{}, fmt.Errorf("Account not found")
}
func accpetAmount() float64 {
	amount := 0.0
	fmt.Print("Enter amount	:	")
	fmt.Scanln(&amount)
	return amount
}

func acceptAccNo() int {
	accNo := 1000
	fmt.Print("Enter Valid account Number	:	")
	fmt.Scanln(&accNo)
	if accNo < 1000 && accNo > count {
		fmt.Print("Invalid Account Number -Please enter a valid	:	")
		fmt.Scanln(&accNo)
	}
	return accNo
}
func menuList() int {
	choice := 0
	fmt.Println("0.Exit")
	fmt.Println("1.Create New Account")
	fmt.Println("2.Deposit")
	fmt.Println("3.Withdraw")
	fmt.Println("4.Print Account details")
	fmt.Print("Enter choice	:	")
	fmt.Scanln(&choice)
	return choice

}

func main() {
	fmt.Println("inside main")
	bankObj := NewBank()
	var acc Account
	var choice = 1
	for 0 != choice {
		switch choice {
		case 1:
			acc = NewAccount()
			bankObj.createAccount(&acc)
			fmt.Println("Print Account Details", acc)
		case 2:
			accNo := acceptAccNo()
			bankObj.deposit(accNo, accpetAmount())
			fmt.Println("Remainging Balance", bankObj.getBalance(accNo))
		case 3:
			accNo := acceptAccNo()
			bankObj.withdraw(accNo, accpetAmount())
			fmt.Println("Remainging Balance", bankObj.getBalance(accNo))
		case 4:
			acc, err := bankObj.printAccountDetails(acceptAccNo())
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Account Details ", acc)
			}

		}
		choice = menuList()
	}

}
