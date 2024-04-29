package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	From      int       `json:"from"`
	To        int       `json:"to"`
	Amount    int       `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}
type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(firstName, lastName, email string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Number:    int64(rand.Intn(100000)),
		CreatedAt: time.Now().UTC(),
	}
}
