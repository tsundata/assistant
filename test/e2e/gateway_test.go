package e2e

import (
	"fmt"
	"github.com/gavv/httpexpect/v2"
	"github.com/go-resty/resty/v2"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/util"
	"log"
	"net/http"
	"testing"
)

const GatewayBaseURL = "http://127.0.0.1:5000"

func TestGatewayIndex(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.GET("/").
		Expect().
		Status(http.StatusOK).Text().Contains("Gateway")
}

// === command ===

func TestTestCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`test`)).
		Expect().
		Status(http.StatusOK)
}

func TestMessageCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`hello world`)).
		Expect().
		Status(http.StatusOK)
}

func TestHelpCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`help`)).
		Expect().
		Status(http.StatusOK).Text().Contains("available commands")
}

func TestVersionCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`version`)).
		Expect().
		Status(http.StatusOK).Text().Contains("Version")
}

func TestQrCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`qr 123456`)).
		Expect().
		Status(http.StatusOK).Text().Contains("http://127.0.0.1:7000/qr/123456")
}

func TestUtCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`ut 1`)).
		Expect().
		Status(http.StatusOK).Text().Contains("1970-01-01 08:00:01 +0800 CST")
}

func TestMenuCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`menu`)).
		Expect().
		Status(http.StatusOK).
		Text().Match(`Memo\nhttp://127.0.0.1:7000/memo/\w{8}-\w{4}-\w{4}-\w{4}-\w{12}`)
}

func TestRandCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`rand 1 100`)).
		Expect().
		Status(http.StatusOK).Text().Match(`\d+`)
}

func TestPwdCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`pwd 32`)).
		Expect().
		Status(http.StatusOK).Text().Length().Equal(32)
}

func TestSubListCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`subs list`)).
		Expect().
		Status(http.StatusOK).Text().Contains("Subscribe:")
}

func TestSubsOpenCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`subs open demo`)).
		Expect().
		Status(http.StatusOK).Text().Contains("ok")
}

func TestSubsCloseCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`subs close demo`)).
		Expect().
		Status(http.StatusOK).Text().Contains("ok")
}

func TestViewCommand(t *testing.T) {
	t.Skip()// fixme
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`view 1`)).
		Expect().
		Status(http.StatusOK).Text().Contains("hello world")
}

func TestRunCommand(t *testing.T) {
	t.Skip()// fixme
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`run 1`)).
		Expect().
		Status(http.StatusOK).Text().Contains("Not running")
}

func TestDocCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`doc`)).
		Expect().
		Status(http.StatusOK).Text().Contains("get [any]")
}

func TestStatsCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`stats`)).
		Expect().
		Status(http.StatusOK).Text().Contains("/pb.Message/Create")
}

func TestTodoCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`todo something`)).
		Expect().
		Status(http.StatusOK).Text().Contains("success")
}

func TestRoleCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`role`)).
		Expect().
		Status(http.StatusOK).Text().Contains("/role/")
}

func TestPinyinCommand(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/debug/event").
		WithBytes([]byte(`pinyin 测试`)).
		Expect().
		Status(http.StatusOK).Text().Contains("ce, shi")
}

// === api ===

func TestAuth(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.POST("/auth").
		WithJSON(pb.TextRequest{Text: getToken()}).
		Expect().Status(http.StatusOK)
}

func TestPage(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.GET("/page").
		WithHeader("Authorization", getAuth()).
		WithJSON(pb.MessageRequest{Uuid: "test"}).
		Expect().Status(http.StatusOK).
		JSON().Object()
}

func TestApps(t *testing.T) {
	e := httpexpect.New(t, GatewayBaseURL)
	e.GET("/apps").
		WithHeader("Authorization", getAuth()).
		Expect().Status(http.StatusOK).
		JSON().Object().ContainsKey("apps")
}

func getToken() string {
	r := resty.New()
	r.SetHostURL(GatewayBaseURL)
	resp, err := r.R().SetBody("menu").Post("/debug/event")
	if err != nil {
		panic(err)
	}
	token := util.ExtractUUID(util.ByteToString(resp.Body()))
	log.Println("Token", token)
	return token
}

func getAuth() string {
	token := getToken()
	return fmt.Sprintf("Bearer %s", token)
}
