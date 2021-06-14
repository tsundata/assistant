package http

import (
	"encoding/json"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/valyala/fasthttp"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) PostJSON(url string, body interface{}) (*fasthttp.Response, error) {
	j, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req := fasthttp.AcquireRequest()
	req.SetBody(j)
	req.Header.SetMethodBytes(util.StringToByte("POST"))
	req.Header.SetContentType("application/json")
	req.SetRequestURIBytes(util.StringToByte(url))
	res := fasthttp.AcquireResponse()
	err = fasthttp.Do(req, res)
	if err != nil {
		return nil, err
	}
	fasthttp.ReleaseRequest(req)

	return res, nil
}
