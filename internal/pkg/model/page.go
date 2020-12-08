package model

import "time"

type Page struct {
	ID      int
	UUID    string
	Title   string
	Content string
	Time    time.Time
}
