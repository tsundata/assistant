package opcode

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"regexp"
	"strings"
)

type Query struct{}

func NewQuery() *Query {
	return &Query{}
}

func (o *Query) Type() int {
	return TypeOp
}

func (o *Query) Doc() string {
	return "query [string:(css|json|regex)] [string] [string]? : (any -> any)"
}

func (o *Query) Run(_ context.Context, inCtx *inside.Context, _ component.Component, params []interface{}) (interface{}, error) {
	if len(params) < 2 {
		return false, app.ErrInvalidParameter
	}
	if t, ok := params[0].(string); ok {
		expression, ok := params[1].(string)
		if !ok {
			return false, errors.New("error param[1] type")
		}

		switch strings.ToLower(t) {
		case "css":
			if len(params) < 3 {
				return false, errors.New("error params[2]")
			}
			attr, ok := params[2].(string)
			if !ok {
				return false, errors.New("error param[2] type")
			}

			if text, ok := inCtx.Value.(string); ok {
				doc, err := goquery.NewDocumentFromReader(strings.NewReader(text))
				if err != nil {
					return nil, err
				}
				elem := doc.Find(expression)
				if elem.Length() > 1 {
					var values []string
					elem.Each(func(i int, s *goquery.Selection) {
						if attr == "text" {
							values = append(values, s.Text())
						} else {
							if v, ex := s.Attr(attr); ex {
								values = append(values, v)
							}
						}
					})
					inCtx.SetValue(values)
					return values, nil
				} else {
					var value string
					if attr == "text" {
						value = elem.Text()
					} else {
						if v, ex := elem.Attr(attr); ex {
							value = v
						}
					}
					inCtx.SetValue(value)
					return value, nil
				}
			}
		case "json":
			if text, ok := inCtx.Value.(string); ok {
				j := gjson.Parse(text)
				value := j.Get(expression)
				result := value.Value()
				inCtx.SetValue(result)
				return result, nil
			}
		case "regex":
			if text, ok := inCtx.Value.(string); ok {
				re, err := regexp.Compile(fmt.Sprintf(`(?m)%s`, expression))
				if err != nil {
					return nil, err
				}
				result := re.FindAllString(text, -1)
				inCtx.SetValue(result)
				return result, nil
			}
		default:
			return false, app.ErrInvalidParameter
		}
	}
	return false, nil
}
