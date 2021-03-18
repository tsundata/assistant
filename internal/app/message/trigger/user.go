package trigger

import "fmt"

type User struct {
	text string
}

func NewUser() *User {
	return &User{}
}

func (t *User) Cond(text string) bool {
	return true
}

func (t *User) Handle() {
	fmt.Println("User handle")
}
