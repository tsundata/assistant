package model

import "time"

type Page struct {
	ID      int
	UUID    string `json:"uuid"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Time    time.Time
}
