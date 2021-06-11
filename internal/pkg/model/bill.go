package model

import "time"

type Bill struct {
	ID          int       `db:"id"`
	Date        string    `db:"date"`
	Payee       string    `db:"payee"`
	Description string    `db:"description"`
	Amount      float64   `db:"amount"`
	Time        time.Time `db:"time"`
}

type BillRecord struct {
	ID      int     `db:"id"`
	BillID  int     `db:"bill_id"`
	Posting string  `db:"posting"` // -> assets
	Amount  float64 `db:"amount"`
}

type Assets struct {
	ID        int       `db:"id"`
	AccountID int       `db:"account_id"` // -> account
	Name      string    `db:"name"`
	Category  string    `db:"category"`
	Balance   float64   `db:"balance"`
	Time      time.Time `db:"time"`
}

type Account struct {
	ID      int       `db:"id"`
	Name    string    `db:"name"`
	Balance float64   `db:"balance"`
	Time    time.Time `db:"time"`
}
