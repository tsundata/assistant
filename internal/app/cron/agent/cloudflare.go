package agent

import (
	"context"
	"encoding/json"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors/cloudflare"
	"time"
)

func DomainAnalyticsReport(b *rulebot.RuleBot) []result.Result {
	// get key
	ctx := context.Background()
	reply, err := rpcclient.GetMiddleClient(b.Client).GetCredential(ctx, &pb.CredentialRequest{Name: cloudflare.ID})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}
	token := ""
	zone := ""
	for _, item := range reply.GetContent() {
		if item.Key == cloudflare.Token {
			token = item.Value
		}
		if item.Key == cloudflare.ZoneID {
			zone = item.Value
		}
	}
	if token == "" || zone == "" {
		return []result.Result{result.EmptyResult()}
	}

	// get dashboard
	end := time.Now().Format("2006-01-02T15:04:05Z")
	start := time.Now().Add(time.Duration(-7*24) * time.Hour).Format("2006-01-02T15:04:05Z")

	cf := cloudflare.NewCloudflare(token, zone)
	analytics, err := cf.GetAnalytics(start, end)
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	j, err := json.Marshal(analytics)
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	return []result.Result{result.MessageResult(util.ByteToString(j))}
}
