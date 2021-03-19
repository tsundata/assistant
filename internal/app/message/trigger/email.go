package trigger

import "fmt"

type Email struct{
	text string
}

func NewEmail() *Email {
	return &Email{}
}

func (t *Email) Cond(text string) bool {
	return true
}

func (t *Email) Handle() {
	fmt.Println("Email handle", t.text)
}
