package model

import "time"

type Credential struct {
	ID      int       `db:"id"`
	Name    string    `db:"name"`
	Type    string    `db:"type"`
	Content string    `db:"content"`
	Time    time.Time `db:"time"`
}
