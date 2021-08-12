package cloudflare

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

const (
	ID        = "cloudflare"
	Token     = "token"
	ZoneID    = "zone_id"
	AccountID = "account_id"
)

type AnalyticResponse struct {
	Data struct {
		Viewer struct {
			Zones []struct {
				FirewallEventsAdaptive []struct {
					Action                string     `json:"action"`
					ClientRequestHTTPHost string     `json:"clientRequestHTTPHost"`
					Datetime              *time.Time `json:"datetime"`
					RayName               string     `json:"rayName"`
					UserAgent             string     `json:"userAgent"`
				} `json:"firewallEventsAdaptive"`
			} `json:"zones"`
		} `json:"viewer"`
	} `json:"data"`
	Errors interface{} `json:"errors"`
}

type Cloudflare struct {
	token  string
	zoneID string
}

func NewCloudflare(token string, zoneID string) *Cloudflare {
	return &Cloudflare{token: token, zoneID: zoneID}
}

func (v *Cloudflare) GetAnalytics(start, end string) (*AnalyticResponse, error) {
	c := resty.New()
	resp, err := c.R().
		SetAuthToken(v.token).
		SetResult(&AnalyticResponse{}).
		SetBody(fmt.Sprintf(`
query
{
  viewer
  {
    zones(filter: { zoneTag: "%s"})
    {
      firewallEventsAdaptive(
          filter: {
            datetime_gt: "%s",
            datetime_lt: "%s" 
          },
          limit: 2,
          orderBy: [datetime_DESC, rayName_DESC])
      {
        action
        datetime
        rayName
        clientRequestHTTPHost
        userAgent
      }
    }
  }
}
`, v.zoneID, start, end)).
		Post("https://api.cloudflare.com/client/v4/graphql")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		result := resp.Result().(*AnalyticResponse)
		return result, nil
	}
	return nil, fmt.Errorf("cloudflare api error %d", resp.StatusCode())
}
