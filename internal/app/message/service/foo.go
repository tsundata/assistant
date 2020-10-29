package service

import (
	"time"

	"github.com/tsundata/assistant/api"
)

type Foo int

func (f Foo) Sum(args api.Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func (f Foo) Sleep(args api.Args, reply *int) error {
	time.Sleep(time.Second * time.Duration(args.Num1))
	*reply = args.Num1 + args.Num2
	return nil
}
