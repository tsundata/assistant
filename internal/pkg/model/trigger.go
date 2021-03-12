package model

import "time"

type Trigger struct {
	ID        int       `db:"id"`
	Type      string    `db:"type"`
	Kind      string    `db:"kind"`
	Flag      string    `db:"flag"`
	Secret    string    `db:"secret"`
	When      string    `db:"when"`
	MessageID int       `db:"message_id"`
	Time      time.Time `db:"time"`
}
