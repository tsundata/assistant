package opcode

import (
	"context"
	"errors"
	"github.com/sourcegraph/checkup/check/dns"
	"github.com/sourcegraph/checkup/check/http"
	"github.com/sourcegraph/checkup/check/tcp"
	"github.com/sourcegraph/checkup/check/tls"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"strings"
)

type Status struct{}

func NewStatus() *Status {
	return &Status{}
}

func (o *Status) Type() int {
	return TypeOp
}

func (o *Status) Doc() string {
	return "status [string:(http|tcp|dns|tls)] [string] : (nil -> bool)"
}

func (o *Status) Run(_ context.Context, _ *inside.Component, params []interface{}) (interface{}, error) {
	if len(params) != 2 {
		return false, app.ErrInvalidParameter
	}
	if checker, ok := params[0].(string); ok {
		text, ok := params[1].(string)
		if !ok {
			return false, errors.New("error param[1]")
		}

		switch strings.ToLower(checker) {
		case "http":
			c := http.Checker{
				Name:     "HTTP check",
				URL:      text,
				Attempts: 3,
			}
			result, err := c.Check()
			if err != nil {
				return false, err
			}
			return result.Healthy, nil
		case "tcp":
			c := tcp.Checker{
				Name:     "TCP check",
				URL:      text,
				Attempts: 3,
			}
			result, err := c.Check()
			if err != nil {
				return false, err
			}
			return result.Healthy, nil
		case "dns":
			c := dns.Checker{
				Name:     "DNS check",
				URL:      text,
				Attempts: 3,
			}
			result, err := c.Check()
			if err != nil {
				return false, err
			}
			return result.Healthy, nil
		case "tls":
			c := tls.Checker{
				Name:     "DNS check",
				URL:      text,
				Attempts: 3,
			}
			result, err := c.Check()
			if err != nil {
				return false, err
			}
			return result.Healthy, nil
		default:
			return false, app.ErrInvalidParameter
		}
	}
	return false, nil
}
