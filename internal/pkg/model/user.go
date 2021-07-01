package model

import "time"

const SuperUserID = 1

type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Mobile    string    `db:"mobile"`
	Remark    string    `db:"remark"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
