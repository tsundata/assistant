package model

import "time"

const (
	RepeatDaily    = "daily"
	RepeatAnnually = "annually"
	RepeatMonthly  = "monthly"
	RepeatWeekly   = "weekly"
	RepeatHourly   = "hourly"
)

type Todo struct {
	ID int64 `db:"id"`

	Content  string `db:"content"`
	Category string `db:"category"`
	Remark   string `db:"remark"`

	Priority int64 `db:"priority"`

	IsRemindAtTime bool       `db:"is_remind_at_time"`
	RemindAt       *time.Time `db:"remind_at"`
	RepeatMethod   string     `db:"repeat_method"`
	RepeatRule     string     `db:"repeat_rule"`
	RepeatEndAt    *time.Time `db:"repeat_end_at"`

	Complete bool `db:"complete"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
