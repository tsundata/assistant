package vendors

import (
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/vendors/dropbox"
	"github.com/tsundata/assistant/internal/pkg/vendors/github"
	"github.com/tsundata/assistant/internal/pkg/vendors/pocket"
	"testing"
)

func TestPocketProvider(t *testing.T) {
	rdb, err := CreateRedisClient(app.Message)
	if err != nil {
		t.Fatal(err)
	}

	p := NewOAuthProvider(rdb, pocket.ID, "https://test/oauth")
	p.AuthorizeURL()
}

func TestGithubProvider(t *testing.T) {
	rdb, err := CreateRedisClient(app.Message)
	if err != nil {
		t.Fatal(err)
	}

	p := NewOAuthProvider(rdb, github.ID, "https://test/oauth")
	p.AuthorizeURL()
}

func TestDropboxProvider(t *testing.T) {
	rdb, err := CreateRedisClient(app.Message)
	if err != nil {
		t.Fatal(err)
	}

	p := NewOAuthProvider(rdb, dropbox.ID, "https://test/oauth")
	p.AuthorizeURL()
}
