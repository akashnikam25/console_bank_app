# Simple Bank Application
This is a simple bank application written in Go programming language. It allows you to create a new account, deposit and withdraw money, and print account details.

## Usage
Clone the repository to your local machine.
Navigate to the cloned directory.
Run go run main.go in the terminal.
Follow the instructions in the terminal to perform various operations.
## Features
Create a new bank account with account name, type, and balance.
Deposit money into an existing account.
Withdraw money from an existing account.
Print account details including account number, name, type, and balance.
Code Explanation
The code defines two structs, Account and Bank. Account contains the fields for account name, account number, account type, and balance. Bank contains an array of Account objects.

The main function creates a new Bank object and initializes a variable called choice to 1. It then enters into a loop that displays a menu of options to the user and waits for the user to enter a choice. The loop continues until the user enters 0 to exit the application.

The options available in the menu are as follows:

- Create a new account: This option allows the user to create a new bank account by entering the account name, type, and balance. The NewAccount function is used to accept user input and return an Account object, which is then added to the Bank object's Accounts array using the createAccount method.
- Deposit: This option allows the user to deposit money into an existing account. The user is prompted to enter the account number and the amount to be deposited. The deposit method of the Bank object is then called to update the balance of the account.
- Withdraw: This option allows the user to withdraw money from an existing account. The user is prompted to enter the account number and the amount to be withdrawn. The withdraw method of the Bank object is then called to update the balance of the account.
- Print account details: This option allows the user to print the details of an existing account. The user is prompted to enter the account number, and the printAccountDetails method of the Bank object is called to return the Account object, which is then printed to the console.

## Conclusion
This simple bank application demonstrates how to use Go programming language to create a command-line application that allows users to create and manage bank accounts. With a little bit of modification, this application can be extended to include more features, such as transferring money between accounts, viewing transaction history, and so on