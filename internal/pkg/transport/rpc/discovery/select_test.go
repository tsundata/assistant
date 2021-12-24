package discovery

import (
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/config"
	"testing"
)

func TestSvcAddr(t *testing.T) {
	c := config.AppConfig{}
	c.SvcAddr.Message = "message:6000"
	c.SvcAddr.Bot = "bot:6005"
	r := SvcAddr(&c, enum.Message)
	require.Equal(t, "message:6000", r)

	r = SvcAddr(&c, enum.Chatbot)
	require.Equal(t, "", r)

	r = SvcAddr(&c, enum.Bot)
	require.Equal(t, "bot:6005", r)

	r = SvcAddr(&c, "NotFound")
	require.Equal(t, "NotFound-error-svc-addr", r)
}
