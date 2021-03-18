package trigger

type Trigger interface {
	Cond(text string) bool
	Handle()
}

func Triggers() []Trigger {
	return []Trigger{
		NewUrl(),
		NewEmail(),
		NewTag(),
		NewUser(),
	}
}
