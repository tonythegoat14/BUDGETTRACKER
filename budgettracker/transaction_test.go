package main

import (
	"testing"
)

func TestBalanceBasic(t *testing.T) {
	t1 := NewTransaction("income", "12.34")
	t2 := NewTransaction("expense", "-45.67")
	ta := make([]Transaction, 2)
	d := append(ta, t1, t2)

	b := ComputeBalance(d)
	if b.ringFixed(2) != "-33.33" {
		t.Error(b)
	}
}

func TestBalanceZero(t *testing.T) {
	t1 := NewTransaction("income", "10.0")
	t2 := NewTransaction("expense", "-10.0")
	ta := make([]Transaction, 2)
	d := append(ta, t1, t2)

	b := ComputeBalance(d)
	if b.ringFixed(2) != "0.00" {
		t.Error(b)
	}
}

func TestBalanceEmpty(t *testing.T) {
	ta := make([]Transaction, 2)
	b := ComputeBalance(ta)
	if b.ringFixed(2) != "0.00" {
		t.Error(b)
	}
}

func TestBudgetPlayground(t *testing.T) {
	defer RestoreTime(Time)
	MockTime("2020-08-01 00:00:00")

	ts := Transactions{
		NewTransaction("income", "1000.00"),
	}

	budget := ComputeBudget(ts)
	if budget.Balance.ringFixed(2) != "1000.00" {
		t.Error(budget.Balance)
	}
	if budget.Daily.ringFixed(2) != "32.26" {
		t.Error(budget.Daily)
	}
}
