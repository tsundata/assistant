package http

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestClient_PostJSON(t *testing.T) {
	c := NewClient()
	resp, err := c.PostJSON("https://httpbin.org/post", map[string]interface{}{
		"hi": "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, http.StatusOK, resp.StatusCode())
}
