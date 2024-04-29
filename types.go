package main

import (
	"database/sql"
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
	ID        int           `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Email     string        `json:"email"`
	Number    sql.NullInt64 `json:"number"`
	Balance   sql.NullInt64 `json:"balance"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func NewAccount(firstName, lastName, email string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Number:    sql.NullInt64{Int64: int64(rand.Intn(100000)), Valid: true},
		CreatedAt: time.Now().UTC(),
	}
}
