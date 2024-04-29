package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(a *Account) error
	DeleteAccount(id int) error
	UpdateAccount(a *Account) error
	GetAccount(id int) (*Account, error)
	GetAccounts() ([]*Account, error)
	GetAccountByEmail(email string) (*Account, error)
	GetAccountByID(id int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password='FishFilet20%' sslmode=disable"
	// docker run --name some-postgres -e POSTGRES_PASSWORD=FishFilet20% -p 5432:5432 -d postgres
	// docker restart some-postgres
	db, error := sql.Open("postgres", connStr)
	if error != nil {
		return nil, error
	}
	if error = db.Ping(); error != nil {
		return nil, error
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreatAcountTable()
}

func (s *PostgresStore) CreatAcountTable() error {
	_, error := s.db.Exec(`CREATE TABLE IF NOT EXISTS accounts (id SERIAL PRIMARY KEY, first_name TEXT, last_name TEXT, email TEXT, number SERIAL, balance SERIAL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`)
	return error
}

func (s *PostgresStore) CreateAccount(a *Account) error {
	_, error := s.db.Exec(`INSERT INTO accounts (first_name, last_name, email) VALUES ($1, $2, $3)`, a.FirstName, a.LastName, a.Email)
	return error
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, error := s.db.Exec(`DELETE FROM accounts WHERE id = $1`, id)
	return error
}

func (s *PostgresStore) UpdateAccount(a *Account) error {
	_, error := s.db.Exec(`UPDATE accounts SET first_name = $1, last_name = $2, email = $3 WHERE id = $4`, a.FirstName, a.LastName, a.Email, a.ID)
	return error
}

func (s *PostgresStore) GetAccount(id int) (*Account, error) {
	var a Account
	error := s.db.QueryRow(`SELECT * FROM accounts WHERE id = $1`, id).Scan(&a.ID, &a.FirstName, &a.LastName, &a.Email)
	if error != nil {
		return nil, error
	}
	return &a, nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, error := s.db.Query(`SELECT * FROM accounts`)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	accounts := []*Account{}
	for rows.Next() {
		account, error := scanIntoAccount(rows)
		if error != nil {
			return nil, error
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (s *PostgresStore) GetAccountByEmail(email string) (*Account, error) {
	var a Account
	error := s.db.QueryRow(`SELECT * FROM accounts WHERE email = $1`, email).Scan(&a.Email)
	if error != nil {
		return nil, error
	}
	return &a, nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	rows, error := s.db.Query(`SELECT * FROM accounts WHERE id = $1`, id)
	if error != nil {
		return nil, error
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	error := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Email,
		&account.Number,
		&account.Balance,
	)
	if error != nil {
		return nil, error
	}
	return account, nil
}

func (s *PostgresStore) Close() {
	s.db.Close()
}
