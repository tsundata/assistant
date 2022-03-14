package vendors

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/tsundata/assistant/api/pb"
	push2 "github.com/tsundata/assistant/internal/pkg/push"
	"github.com/tsundata/assistant/internal/pkg/vendors/bark"
	"github.com/tsundata/assistant/internal/pkg/vendors/cloudcone"
	"github.com/tsundata/assistant/internal/pkg/vendors/cloudflare"
	"github.com/tsundata/assistant/internal/pkg/vendors/doctorxiong"
	"github.com/tsundata/assistant/internal/pkg/vendors/dropbox"
	"github.com/tsundata/assistant/internal/pkg/vendors/email"
	"github.com/tsundata/assistant/internal/pkg/vendors/github"
	"github.com/tsundata/assistant/internal/pkg/vendors/pocket"
	"github.com/tsundata/assistant/internal/pkg/vendors/pushover"
)

var OAuthProviderApps = []string{
	github.ID,
	pocket.ID,
	dropbox.ID,
}

var ProviderCredentialOptions = map[string]interface{}{
	// OAuth
	github.ID: map[string]string{
		github.ClientIdKey:     "Client ID",
		github.ClientSecretKey: "Client secrets",
	},
	pocket.ID: map[string]string{
		pocket.ClientIdKey: "Consumer Key",
	},
	dropbox.ID: map[string]string{
		dropbox.ClientIdKey:     "App key",
		dropbox.ClientSecretKey: "App secret",
	},

	// Push
	pushover.ID: map[string]string{
		pushover.TokenKey: "App Token",
		pushover.UserKey:  "User Key",
	},
	bark.ID: map[string]string{
		bark.ServerUrl: "Server URL",
		bark.TokenKey:  "Token",
	},

	// Service
	email.ID: map[string]string{
		email.Host:     "SMTP Host",
		email.Port:     "SMTP Port",
		email.Username: "Username Mail",
		email.Password: "Password",
	},

	// API
	cloudflare.ID: map[string]string{
		cloudflare.Token:     "Api Token",
		cloudflare.ZoneID:    "Zone ID",
		cloudflare.AccountID: "Account ID",
	},
	cloudcone.ID: map[string]string{
		cloudcone.ApiKey:  "Api Key",
		cloudcone.ApiHash: "Api Hash",
	},
	doctorxiong.ID: map[string]string{
		doctorxiong.Token: "Api Token",
	},
}

type OAuthProvider interface {
	AuthorizeURL() string
	GetAccessToken(code string) (interface{}, error)
	Redirect(c *fiber.Ctx, middle pb.MiddleSvcClient) error
	StoreAccessToken(c *fiber.Ctx, middle pb.MiddleSvcClient) error
}

func NewOAuthProvider(rdb *redis.Client, category, url string) OAuthProvider {
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

type PushProvider interface {
	Send(message push2.Message) error
}
