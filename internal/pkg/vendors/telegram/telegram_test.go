package telegram

import (
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTelegram_SendMessage(t *testing.T) {
	token := "test"
	tg := NewTelegram(token)

	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	tg.c = client

	fixture := `{
    "ok": true,
    "result": {}
}`
	responder := httpmock.NewStringResponder(200, fixture)
	//fakeUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	fakeUrl := "//%2FsendMessage/sendMessage"
	httpmock.RegisterResponder("POST", fakeUrl, responder)

	res, err := tg.SendMessage(1, "hi")
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, res.Ok)

	httpmock.DeactivateAndReset()
}
