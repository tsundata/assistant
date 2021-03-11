package vendors

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/vendors/dropbox"
	"github.com/tsundata/assistant/internal/pkg/vendors/github"
	"github.com/tsundata/assistant/internal/pkg/vendors/pocket"
	"github.com/tsundata/assistant/internal/pkg/vendors/pushover"
)

var OAuthProviders = []string{
	github.ID,
	pocket.ID,
	pushover.ID,
	dropbox.ID,
}

var OAuthProvidersOptions = map[string]interface{}{
	github.ID: map[string]string{
		github.ClientIdKey:     "Client ID",
		github.ClientSecretKey: "Client secrets",
	},
	pocket.ID: map[string]string{
		pocket.ClientIdKey: "Consumer Key",
	},
	pushover.ID: map[string]string{
		pushover.TokenKey: "API Token",
		pushover.UserKey:  "User Key",
	},
	dropbox.ID: map[string]string{
		dropbox.ClientIdKey:     "App key",
		dropbox.ClientSecretKey: "App secret",
	},
}

type OAuthProvider interface {
	AuthorizeURL() string
	GetAccessToken(code string) (interface{}, error)
	Redirect(c *fiber.Ctx, mid pb.MiddleClient) error
	StoreAccessToken(c *fiber.Ctx, mid pb.MiddleClient) error
}

func NewOAuthProvider(rdb *redis.Client, c *fiber.Ctx, url string) OAuthProvider {
	category := c.Params("category")
	redirectURI := fmt.Sprintf("%s/oauth/%s", url, category)
	var provider OAuthProvider

	switch category {
	case pocket.ID:
		p := pocket.NewPocket("", "", redirectURI, "")
		p.SetRDB(rdb)
		provider = p
	case github.ID:
		provider = github.NewGithub("", "", redirectURI, "")
	case dropbox.ID:
		provider = dropbox.NewDropbox("", "", redirectURI, "")
	default:
		return nil
	}

	return provider
}
