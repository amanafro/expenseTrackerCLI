package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	ID         int
	Number     string
	Balance    int
	CustomerID int
}

type Customer struct {
	ID       int
	Name     string
	Password string
}

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gobank")
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	for {
		intro()
	}
}

func intro() {
	var choice string

	fmt.Println("Hello! How can help you? \n 1. Deposit \n 2. Withdraw \n 3. Balance")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		Deposit()
	case "2":
		Withdraw()
	case "3":
		CheckBalance()
	default:
		fmt.Println("Invalid choice")
	}
}

func Deposit() {
	db := InitDB()

	var depositeMoney int

	fmt.Println("Great! How much you wanna deposit?")
	if _, err := fmt.Scanln(&depositeMoney); err != nil {
		log.Fatal("Error while depositing", err)
	}

	if depositeMoney < 5 {
		fmt.Println("You have to atleast deposit CHF 5")
	}

	res, err := db.Exec("UPDATE accounts SET balance = balance + ? WHERE account_id=1 ", depositeMoney)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	db.Close()
}

func Withdraw() {
	db := InitDB()

	var withdrawMoney int
	fmt.Println("How much money would like to Withdraw? ")
	if _, err := fmt.Scanln(&withdrawMoney); err != nil {
		log.Fatal("Error while withdrawing the money")
	}

	if withdrawMoney < 10 {
		fmt.Println("You can't withdraw less than CHF 10")
	}

	update, err := db.Exec("UPDATE accounts SET balance = balance - ? WHERE account_id=1", withdrawMoney)
	if err != nil {
		log.Fatal(err)
	} else {

	}

	fmt.Println(update)
	db.Close()
}

func CheckBalance() {
	db := InitDB()

	update, err := db.Query("SELECT account_id, account_number, balance FROM accounts WHERE customer_id=1 ")
	if err != nil {
		log.Fatal(err)
	}

	for update.Next() {

		var accountBalance Account

		err = update.Scan(&accountBalance.ID, &accountBalance.Number, &accountBalance.Balance)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Account ID: %d,\n Account Num: %s,\n Account Balance: %d\n",
			accountBalance.ID, accountBalance.Number, accountBalance.Balance)
	}

	fmt.Println(update)
	db.Close()
}
