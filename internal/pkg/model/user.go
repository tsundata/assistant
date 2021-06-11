package model

const SuperUserID = 1

type User struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Mobile string `db:"mobile"`
	Remark string `db:"remark"`
	Time   string `db:"time"`
}
