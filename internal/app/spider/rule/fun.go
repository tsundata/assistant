package rule

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/tsundata/assistant/internal/pkg/util"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	rxFunName = regexp.MustCompile(`^[a-z$][a-zA-Z]{0,15}`)
)

func PowerfulFind(s *goquery.Selection, q string) *goquery.Selection {
	rxSelectPseudoEq := regexp.MustCompile(`:eq\(\d+\)`)
	if rxSelectPseudoEq.MatchString(q) {
		rs := rxSelectPseudoEq.FindAllStringIndex(q, -1)
		sel := s
		for _, r := range rs {
			iStr := q[r[0]+4 : r[1]-1]
			i64, _ := strconv.ParseInt(iStr, 10, 32)
			i := int(i64)
			sq := q[:r[0]]
			q = strings.TrimSpace(q[r[1]:])
			sel = sel.Find(sq).Eq(i)
		}
		if len(q) > 0 {
			sel = sel.Find(q)
		}
		return sel
	} else {
		return s.Find(q)
	}
}

type Fun struct {
	Name   string
	Raw    string
	Params []string

	Document  *goquery.Document
	Selection *goquery.Selection
	Result    string

	PrevFun *Fun
	NextFun *Fun
}

func (f *Fun) InitSelector() error {
	if len(f.Params) > 0 {
		f.Selection = PowerfulFind(f.Selection, f.Params[0])
	}
	return nil
}

func (f *Fun) Invoke() (string, error) {
	var err error
	switch f.Name {
	case "$":
		err = f.InitSelector()
	case "attr":
		f.Result, _ = f.PrevFun.Selection.Attr(f.Params[0])
	case "text":
		f.Result = f.PrevFun.Selection.Text()
	case "html":
		f.Result, err = f.PrevFun.Selection.Html()
	case "outerHTML":
		f.Result, err = goquery.OuterHtml(f.PrevFun.Selection)
	case "style":
		f.Result, _ = f.PrevFun.Selection.Attr("style")
	case "href":
		f.Result, _ = f.PrevFun.Selection.Attr("href")
	case "src":
		f.Result, _ = f.PrevFun.Selection.Attr("src")
	case "class":
		f.Result, _ = f.PrevFun.Selection.Attr("class")
	case "id":
		f.Result, _ = f.PrevFun.Selection.Attr("id")
	case "expand":
		rx, err := regexp.Compile(f.Params[0])
		if err != nil {
			return "", err
		}
		src := f.PrevFun.Result
		var dst []byte
		m := rx.FindStringSubmatchIndex(src)
		s := rx.ExpandString(dst, f.Params[1], src, m)
		f.Result = util.ByteToString(s)
	case "match":
		rx, err := regexp.Compile(f.Params[0])
		if err != nil {
			return "", err
		}
		rs := rx.FindAllStringSubmatch(f.PrevFun.Result, -1)
		if len(rs) > 0 && len(rs[0]) > 1 {
			f.Result = rs[0][1]
		}
	}
	if err != nil {
		return "", err
	}
	if f.NextFun != nil {
		return f.NextFun.Invoke()
	} else {
		return f.Result, nil
	}
}

func (f *Fun) Append(s string) (*Fun, *Fun) {
	f.NextFun = ParseFun(f.Selection, s)
	f.NextFun.PrevFun = f
	return f, f.NextFun
}

func ParseFun(sel *goquery.Selection, str string) *Fun {
	fun := new(Fun)
	fun.Raw = str
	fun.Selection = sel

	sa := rxFunName.FindAllString(str, -1)
	fun.Name = sa[0]
	ls := str[len(sa[0]):]
	var ps []string
	p, pl := parseParams(ls)
	for i := 0; ; i++ {
		if v, ok := p["$"+strconv.Itoa(i)]; ok {
			ps = append(ps, v)
		} else {
			break
		}
	}
	if len(ps) > 0 {
		fun.Params = ps
	}
	ls = ls[pl+1:]
	if len(ls) > 0 {
		ls = ls[1:]
		fun.Append(ls)
	}

	return fun
}

// start with "(", will return params map and end pos.
// all params string type:
// (key1 = 0, key2 = "str_exam\"ple", key3 = `exp_\`example\n`)
// (key1 = 0, key2, key3)
// (key1, key2, key3)
// ("str_exam\"ple", /exp_\/example\n/, 2)
// source: https://github.com/wspl/creeper/blob/eb1753da1c54ade30e8e6ee82e1923b4473dbc13/town.go
func parseParams(s string) (map[string]string, int) {
	endPos := -1

	kvMap := map[string]string{}
	pK := ""
	pIsK := false

	var sb bytes.Buffer

	inKey := false
	inStr := false // "example"
	inExp := false // `example`
	inStd := false //  example

	noKeyIndex := 0
	insertVal := func(v string) {
		if pIsK {
			kvMap[pK] = strings.TrimSpace(sb.String())
		} else {
			kvMap["$"+strconv.Itoa(noKeyIndex)] = v
			noKeyIndex++
		}
		pIsK = false
	}

	for i, c := range s {
		cso := func(o int) int32 {
			oi := i + o
			if oi >= 0 && oi < len(s) {
				return rune(s[oi])
			}
			return 0
		}
		co := func(o int) int32 {
			if i+o < 0 || i+0 >= len(s) {
				return 0
			}
			if o < 0 {
				j := i
				for j >= 0 && o != 0 {
					j--
					if !unicode.IsSpace(rune(s[j])) {
						o++
					}

				}
				return rune(s[j])
			} else if o > 0 {
				j := i
				for j < len(s)-1 && o != 0 {
					j++
					if !unicode.IsSpace(rune(s[j])) {
						o--
					}
				}
				return rune(s[j])
			} else {
				return rune(s[i])
			}
		}

		if i == 0 && c != '(' {
			return nil, -1
		}

		if !inExp && !inStr && !inStd {
			if (co(-1) == '(' || co(-1) == ',') && (unicode.IsLetter(c) || c == '@') {
				inKey = true
			} else if (co(-1) == '=' || co(-1) == ',' || co(-1) == '(') &&
				!unicode.IsSpace(c) && c != '"' && c != '`' {
				inStd = true
			} else if co(-2) == '=' || co(-2) == ',' || co(-2) == '(' {
				switch co(-1) {
				case '"':
					inStr = true
				case '`':
					inExp = true
				}
			}
		}

		if inKey || inExp || inStd || inStr {
			sb.WriteRune(c)
		}

		if !inExp && !inStd && !inStr && !inKey && c == ')' {
			endPos = i
		}

		if c != '\\' {
			if inKey && (co(1) == ',' || co(1) == ')' || co(1) == '=') {
				inKey = false
				pK = strings.TrimSpace(sb.String())
				kvMap[pK] = ""
				if co(1) != ',' {
					pIsK = true
				}
				sb.Reset()
			} else if inStr && cso(1) == '"' {
				inStr = false
				s := strings.TrimSpace(sb.String())
				s = strings.ReplaceAll(s, `\\`, `\`)
				insertVal(s)
				sb.Reset()
			} else if inExp && cso(1) == '`' {
				inExp = false
				s := strings.TrimSpace(sb.String())
				insertVal(s)
				sb.Reset()
			} else if inStd && (co(1) == ',' || co(1) == ')') {
				inStd = false
				s := strings.TrimSpace(sb.String())
				insertVal(s)
				sb.Reset()
			}
		}

		if endPos > -1 {
			break
		}
	}

	return kvMap, endPos
}
