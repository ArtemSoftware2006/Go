package main

import "errors"

type Account struct {
	owner   string
	balance float64
}

func NewAccount(owner string) *Account {
	return &Account{
		owner:   owner,
		balance: 0,
	}
}

func (account *Account) SetBalance(amount float64) error {
	if amount < 0 {
		return errors.New("нельзя установить отрицательный баланс")
	}
	account.balance = amount
	return nil
}

func (account *Account) GetBalance() float64 {
	return account.balance
}

func (account *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("нельзя положить отрицательную сумму")
	}
	account.balance += amount
	return nil
}

func (account *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("нельзя снять отрицательную сумму")
	} else if account.balance < amount {
		return errors.New("недостаточно средств")
	}
	account.balance -= amount
	return nil
}
