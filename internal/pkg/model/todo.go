package model

import "time"

type Todo struct {
	ID             int        `db:"id"`
	Content        string     `db:"content"`
	Priority       int        `db:"priority"`
	IsRemindAtTime bool       `db:"is_remind_at_time"`
	RemindAt       *time.Time `db:"remind_at"`
	RepeatMethod   string     `db:"repeat_method"`
	RepeatRule     string     `db:"repeat_rule"`
	Category       string     `db:"category"`
	Remark         string     `db:"remark"`
	Complete       bool       `db:"complete"`
	Time           time.Time  `db:"time"`
}