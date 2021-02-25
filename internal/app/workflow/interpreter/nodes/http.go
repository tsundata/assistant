package nodes

import (
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
	name        string
	properties  map[string]interface{}
	credentials map[string]interface{}
}

func (n *HttpNode) Execute(input interface{}) (interface{}, error) {
	method := extractProperties(input, n.properties, "method").(string)
	url := extractProperties(input, n.properties, "url").(string)
	responseFormat := extractProperties(input, n.properties, "response_format").(string)
	extract := extractProperties(input, n.properties, "extract").(map[string]interface{})
	headers := make(map[string]interface{})
	if _, ok := n.properties["headers"]; ok {
		headers = extractProperties(input, n.properties, "headers").(map[string]interface{})
	}
	query := make(map[string]interface{})
	if _, ok := n.properties["query"]; ok {
		query = extractProperties(input, n.properties, "query").(map[string]interface{})
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	var body io.Reader
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// set header
	for k, v := range headers {
		req.Header.Add(k, v.(string))
	}
	// set query
	for k, v := range query {
		req.URL.Query().Set(k, v.(string))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Process
	var result []map[string]interface{}
	switch responseFormat {
	case "html", "xml":
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return nil, err
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

		multipleResult(multiple, htmlResult, &result)
	case "json":
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		jsonText := utils.ByteToString(data)
		gjsonResult := gjson.Parse(jsonText)

		jsonResult := make(map[string]interface{})
		multiple := false
		for field, i := range extract {
			opt := i.(map[string]interface{})
			path := opt["path"].(string)
			valOpt := opt["value"].(string)
			value := gjsonResult.Get(path)
			if value.IsArray() {
				multiple = true
				var values []interface{}
				for _, i := range value.Array() {
					if valOpt == "string(.)" {
						values = append(values, i.Value())
					}
				}
				jsonResult[field] = values
			} else {
				jsonResult[field] = value.Value()
			}
		}

		multipleResult(multiple, jsonResult, &result)
	case "text":
		text, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		textResult := make(map[string]interface{})
		multiple := false
		for field, i := range extract {
			opt := i.(map[string]interface{})
			reg := opt["regexp"].(string)
			valOpt := opt["value"].(string)
			index := opt["index"].(float64)

			re := regexp.MustCompile(reg)
			findResult := re.FindAll(text, int(index))
			if len(findResult) > 1 {
				multiple = true
				var values []interface{}
				for _, i := range findResult {
					if valOpt == "string(.)" {
						values = append(values, utils.ByteToString(i))
					}
				}
				textResult[field] = values
			} else if len(findResult) == 1 {
				textResult[field] = utils.ByteToString(findResult[0])
			} else {
				textResult[field] = ""
			}
		}

		multipleResult(multiple, textResult, &result)
	default:
		return nil, errors.New("response format error: " + responseFormat)
	}

	return result, nil
}

func multipleResult(is bool, data map[string]interface{}, result *[]map[string]interface{}) {
	if is {
		var multipleResult []map[string]interface{}
		length := 0
		for _, v := range data {
			if items, ok := v.([]interface{}); ok {
				length = len(items)
			}
		}

		for i := 0; i < length; i++ {
			r := make(map[string]interface{})
			for k, v := range data {
				if items, ok := v.([]interface{}); ok {
					r[k] = items[i]
				}
			}
			multipleResult = append(multipleResult, r)
		}
		*result = append(*result, multipleResult...)
	} else {
		*result = append(*result, data)
	}
}
