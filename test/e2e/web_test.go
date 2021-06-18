package e2e

import (
	"fmt"
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"
)

const WebBaseURL = "http://127.0.0.1:7000"

func TestWebIndex(t *testing.T) {
	e := httpexpect.New(t, WebBaseURL)
	e.GET("/").
		Expect().
		Status(http.StatusOK).Text().Contains("Web")
}

func TestEchoIndex(t *testing.T) {
	e := httpexpect.New(t, WebBaseURL)
	e.GET("/echo").
		WithQuery("text", "test").
		Expect().
		Status(http.StatusOK).Text().Contains("test")
}

func TestRobotsIndex(t *testing.T) {
	e := httpexpect.New(t, WebBaseURL)
	e.GET("/Robots.txt").
		Expect().
		Status(http.StatusOK).Text().Contains(`User-agent: *
Disallow: /`)
}

func TestGetPageIndex(t *testing.T) {
	e := httpexpect.New(t, WebBaseURL)
	e.GET("/page/7e5833e3-3a55-4228-9775-ce90794897f2").
		Expect().
		Status(http.StatusOK).Body().Contains("test")
}

func TestAppGithubIndex(t *testing.T) {
	e := httpexpect.New(t, WebBaseURL)
	e.GET("/app/github").
		Expect().
		Status(http.StatusOK).Body().Contains("html")
}

func TestOauthGithubIndex(t *testing.T) {
	e := httpexpect.New(t, WebBaseURL)
	e.GET("/oauth/github").
		Expect().
		Status(http.StatusOK).Body().Contains("html")
}

func TestAppDropboxIndex(t *testing.T) {
	e := httpexpect.New(t, WebBaseURL)
	e.GET("/app/dropbox").
		Expect().
		Status(http.StatusOK).Body().Contains("html")
}

func TestOauthDropboxIndex(t *testing.T) {
	e := httpexpect.New(t, WebBaseURL)
	e.GET("/oauth/dropbox").
		Expect().
		Status(http.StatusOK).Body().Contains("html")
}

func TestAppPocketIndex(t *testing.T) {
	e := httpexpect.New(t, WebBaseURL)
	e.GET("/app/pocket").
		Expect().
		Status(http.StatusOK).Body().Contains("html")
}

func TestOauthPocketIndex(t *testing.T) {
	e := httpexpect.New(t, WebBaseURL)
	e.GET("/oauth/pocket").
		Expect().
		Status(http.StatusOK).Body().Contains("html")
}

func TestQrIndex(t *testing.T) {
	e := httpexpect.New(t, WebBaseURL)
	e.GET("/qr/123456").
		Expect().
		Status(http.StatusOK).Body().Contains("PNG")
}

func TestGetMemoIndex(t *testing.T) {
	token := getToken()
	e := httpexpect.New(t, WebBaseURL)
	e.GET(fmt.Sprintf("/memo/%s", token)).
		Expect().
		Status(http.StatusOK).Body().Contains("Memo")
}

func TestWebhookIndex(t *testing.T) {
	e := httpexpect.New(t, WebBaseURL)
	e.GET("/webhook/123456").
		Expect().
		Status(http.StatusOK).Text().Contains("ok")
	e.POST("/webhook/123456").
		Expect().
		Status(http.StatusOK).Text().Contains("ok")
}
