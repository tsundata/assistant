package trigger

import "fmt"

type Url struct{}

func NewUrl() *Url {
	return &Url{}
}

func (t *Url) Cond(text string) bool {
	return true
}

func (t *Url) Handle() {
	fmt.Println("url handle")
}