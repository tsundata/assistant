package nodes

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type HttpNode struct {
	name string
}

func (n HttpNode) Execute(properties map[string]interface{}, credentials map[string]interface{}, input string) (string, error) {
	method := properties["method"].(string)
	url := properties["url"].(string)
	responseFormat := properties["response_format"].(string)
	headers := properties["headers"].(map[string]interface{})
	query := properties["query"].(map[string]interface{})
	extract := properties["extract"].(map[string]interface{})

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	var body io.Reader
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}

	for k, v := range headers {
		req.Header.Add(k, v.(string))
	}

	for k, v := range query {
		req.URL.Query().Set(k, v.(string))
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Process
	var result []map[string]interface{}
	switch responseFormat {
	case "html", "xml":
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return "", err
		}

		htmlResult := make(map[string]interface{})
		multiple := false

		for field, i := range extract {
			opt := i.(map[string]interface{})
			elem := doc.Find(opt["css"].(string))
			attrOpt := opt["value"].(string)

			if elem.Length() > 1 {
				multiple = true
				var values []interface{}
				elem.Each(func(i int, s *goquery.Selection) {
					if attrOpt == "." {
						values = append(values, s.Text())
					} else {
						attr := strings.Replace(attrOpt, "@", "", 1)
						if v, ex := s.Attr(attr); ex {
							values = append(values, v)
						}
					}
				})
				htmlResult[field] = values
			} else {
				if attrOpt == "." {
					htmlResult[field] = elem.Text()
				} else {
					value := strings.Replace(attrOpt, "@", "", 1)
					if v, ex := elem.Attr(value); ex {
						htmlResult[field] = v
					} else {
						htmlResult[field] = ""
					}
				}
			}
		}
		if multiple {
			var multipleResult []map[string]interface{}
			length := 0
			for _, v := range htmlResult {
				if items, ok := v.([]interface{}); ok {
					length = len(items)
				}
			}

			for i := 0; i < length; i++ {
				r := make(map[string]interface{})
				for k, v := range htmlResult {
					if items, ok := v.([]interface{}); ok {
						r[k] = items[i]
					}
				}
				multipleResult = append(multipleResult, r)
			}
			result = multipleResult
		} else {
			result = append(result, htmlResult)
		}
	case "json":
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		jsonText := string(body)
		gjsonResult := gjson.Parse(jsonText)

		jsonResult := make(map[string]interface{})
		multiple := false
		for field, i := range extract {
			opt := i.(map[string]interface{})
			path := opt["path"].(string)
			value := gjsonResult.Get(path)
			if value.IsArray() {
				multiple = true
				var values []interface{}
				for _, i := range value.Array() {
					values = append(values, i.Value())
				}
				jsonResult[field] = values
			} else {
				jsonResult[field] = value.Value()
			}
		}
		if multiple {
			var multipleResult []map[string]interface{}
			length := 0
			for _, v := range jsonResult {
				if items, ok := v.([]interface{}); ok {
					length = len(items)
				}
			}

			for i := 0; i < length; i++ {
				r := make(map[string]interface{})
				for k, v := range jsonResult {
					if items, ok := v.([]interface{}); ok {
						r[k] = items[i]
					}
				}
				multipleResult = append(multipleResult, r)
			}
			result = multipleResult
		} else {
			result = append(result, jsonResult)
		}
	case "text":
		text, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		textResult := make(map[string]interface{})
		multiple := false
		for field, i := range extract {
			opt := i.(map[string]interface{})
			reg := opt["regexp"].(string)
			index := opt["index"].(float64)

			re := regexp.MustCompile(reg)
			findResult := re.FindAll(text, int(index))
			if len(findResult) > 1 {
				multiple = true
				var values []interface{}
				for _, i := range findResult {
					values = append(values, string(i))
				}
				textResult[field] = values
			} else if len(findResult) == 1 {
				textResult[field] = string(findResult[0])
			} else {
				textResult[field] = ""
			}
		}
		if multiple {
			var multipleResult []map[string]interface{}
			length := 0
			for _, v := range textResult {
				if items, ok := v.([]interface{}); ok {
					length = len(items)
				}
			}

			for i := 0; i < length; i++ {
				r := make(map[string]interface{})
				for k, v := range textResult {
					if items, ok := v.([]interface{}); ok {
						r[k] = items[i]
					}
				}
				multipleResult = append(multipleResult, r)
			}
			result = multipleResult
		} else {
			result = append(result, textResult)
		}
	default:
		return "", errors.New("response format error: " + responseFormat)
	}

	d, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return utils.ByteToString(d), nil
}
