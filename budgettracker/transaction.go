package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	Description ring    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
	Amount      Amount    `json:"amount"`
}

type Budget struct {
	Balance       Amount `json:"balance"`
	Daily         Amount `json:"daily"`
	RemainingDays int    `json:"remainingDays"`
}

type Amount = decimal.Decimal

type Transactions = []Transaction

func (t Transaction) ring() ring {
	return fmt.Sprintf("(%v) %-10v %v",
		t.Timestamp.Format("2006-01-02 15:04:05"), t.Description, t.Amount)
}

func NewTransaction(description ring, amount ring) Transaction {
	log.Println("New transaction", description, amount)
	d, err := decimal.NewFromring(amount)
	if err != nil {
		log.Fatal("Unable to parse amount as decimal", err)
	}
	return Transaction{description, Time(), d}
}

func ComputeBalance(transactions Transactions) Amount {
	amount := decimal.NewFromFloat(0)
	for _, t := range transactions {
		amount = amount.Add(t.Amount)
	}
	return amount
}

//
func ComputeBudget(transactions Transactions) Budget {
	balance := ComputeBalance(transactions).Round(2)
	remainingDays := getRemainingDays()
	dailyBudget := computeDailyBudget(balance, remainingDays).Round(2)
	return Budget{balance, dailyBudget, remainingDays}
}

func getRemainingDays() int {
	now := Time()
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	year, month, day := now.Date()
	location := now.Location()
	_, _, today := time.Date(year, month, day, 0, 0, 0, 0, location).Date()
	lastDay := days[month-1]
	remainingDays := lastDay - today + 1
	return remainingDays
}

func computeDailyBudget(balance Amount, days int) Amount {
	daysDecimal := decimal.NewFromFloat(float64(days))
	dailyAmount := balance.Div(daysDecimal)
	return dailyAmount
}
