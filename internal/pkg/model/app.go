package model

import "time"

type App struct {
	ID    int       `db:"id"`
	Name  string    `db:"name"`
	Type  string    `db:"type"`
	Token string    `db:"token"`
	Extra string    `db:"extra"`
	Time  time.Time `db:"time"`
}
