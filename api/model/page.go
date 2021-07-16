package model

import "time"

type Page struct {
	ID        int       `db:"id"`
	UUID      string    `db:"uuid"`
	Type      string    `db:"type"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
