package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
)

const (
	metricsPath = "/metrics"
	faviconPath = "/favicon.ico"
)

var (
	httpHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http_server",
		Name:      "requests_seconds",
		Help:      "Histogram of response latency (seconds) of http handlers.",
	}, []string{"method", "code", "uri"})
)

func init() {
	prometheus.MustRegister(httpHistogram)
}

// GinPrometheus  struct
type FasthttpPrometheus struct {
	start   time.Time
	ignored map[string]bool
}

// Option
type Option func(*FasthttpPrometheus)

// Ignore ignore path
func Ignore(path ...string) Option {
	return func(gp *FasthttpPrometheus) {
		for _, p := range path {
			gp.ignored[p] = true
		}
	}
}

func New(options ...Option) *FasthttpPrometheus {
	gp := &FasthttpPrometheus{
		ignored: map[string]bool{
			metricsPath: true,
			faviconPath: true,
		},
	}
	for _, o := range options {
		o(gp)
	}
	return gp
}

func (fp *FasthttpPrometheus) Start() {
	fp.start = time.Now()
}

func (fp *FasthttpPrometheus) End(ctx *fasthttp.RequestCtx) {
	if fp.ignored[string(ctx.Path())] == true {
		return
	}

	httpHistogram.WithLabelValues(
		string(ctx.Method()),
		strconv.Itoa(ctx.Response.StatusCode()),
		string(ctx.Path()),
	).Observe(time.Since(fp.start).Seconds())
}
