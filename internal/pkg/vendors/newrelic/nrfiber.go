package newrelic

import (
	"github.com/gofiber/fiber/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/valyala/fasthttp"
	"net/http"
	"net/url"
)

// reference: https://github.com/nobuyo/nrfiber

// MiddlewareConfig defines the config for middleware.
type MiddlewareConfig struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// NewRelicApp is newrelic.Application
	//
	// Required. Default: empty Application
	NewRelicApp *newrelic.Application
}

// ConfigDefault is the default config
var ConfigDefault = MiddlewareConfig{
	Next:        nil,
	NewRelicApp: &newrelic.Application{},
}

func transformRequestHeaders(r *fasthttp.Request) http.Header {
	header := make(http.Header)
	r.Header.VisitAll(func(k, v []byte) {
		sk := string(k)
		sv := string(v)
		header.Set(sk, sv)
	})

	return header
}

func transformResponseHeaders(r *fasthttp.Response) http.Header {
	header := make(http.Header)
	r.Header.VisitAll(func(k, v []byte) {
		sk := string(k)
		sv := string(v)
		header.Set(sk, sv)
	})

	return header
}

func toHTTPRequest(ctx *fasthttp.RequestCtx) *http.Request {
	uri := ctx.Request.URI()
	u := &url.URL{
		Scheme:   string(uri.Scheme()),
		Path:     string(uri.Path()),
		Host:     string(uri.Host()),
		RawQuery: string(uri.QueryString()),
	}

	return &http.Request{
		Method: string(ctx.Request.Header.Method()),
		URL:    u,
		Proto:  "HTTP/1.1",
		Header: transformRequestHeaders(&ctx.Request),
		Host:   string(uri.Host()),
		TLS:    ctx.TLSConnectionState(),
	}
}

// NewMiddleware creates a new middleware handler
func NewMiddleware(config ...MiddlewareConfig) fiber.Handler {
	// Set default config
	cfg := ConfigDefault
	// Override config if provided
	if len(config) > 0 {
		cfg = config[0]
	}

	app := cfg.NewRelicApp

	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		txn := app.StartTransaction(c.Path())
		txn.SetWebRequestHTTP(toHTTPRequest(c.Context()))

		c.Locals("newrelic_transaction", txn)

		defer func() {
			rw := txn.SetWebResponse(&ResponseWriter{
				header: transformResponseHeaders(&c.Context().Response),
			})

			rw.WriteHeader(c.Context().Response.StatusCode())

			txn.End()
		}()

		return c.Next()
	}
}

// ResponseWriter imitates http.ResponseWriter
type ResponseWriter struct {
	statusCode int
	header     http.Header
	body       []byte
}

// Header implementation
func (rw *ResponseWriter) Header() http.Header {
	if rw.header == nil {
		rw.header = make(http.Header)
	}
	return rw.header
}

// WriteHeader implementation
func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
}

// Write implementation
func (rw *ResponseWriter) Write(p []byte) (int, error) {
	rw.body = append(rw.body, p...)
	return len(p), nil
}
