package agent

import (
	"context"
	"encoding/json"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/vendors/cloudcone"
	"strconv"
	"strings"
	"time"
)

func CloudconeWeeklyBilling(ctx context.Context, comp component.Component) []result.Result {
	if comp.Middle() == nil {
		return []result.Result{result.EmptyResult()}
	}
	// get key
	reply, err := comp.Middle().GetCredential(ctx, &pb.CredentialRequest{Type: cloudcone.ID})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}
	key := ""
	hash := ""
	for _, item := range reply.GetContent() {
		if item.Key == cloudcone.ApiKey {
			key = item.Value
		}
		if item.Key == cloudcone.ApiHash {
			hash = item.Value
		}
	}
	if key == "" || hash == "" {
		return []result.Result{result.EmptyResult()}
	}

	// get
	cn := cloudcone.NewCloudCone(key, hash)
	instances, err := cn.GetInstances()
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	var res strings.Builder
	res.WriteString("CloudCone Billing (")
	res.WriteString(time.Now().Format("2006-01-02"))
	res.WriteString(")\n\n")
	res.WriteString("Total Instances: ")
	res.WriteString(strconv.Itoa(instances.Data.TotalInstances))
	res.WriteString("\n")
	res.WriteString("Total CPU: ")
	res.WriteString(strconv.Itoa(instances.Data.TotalCpu))
	res.WriteString("\n")
	res.WriteString("Total RAM: ")
	res.WriteString(strconv.Itoa(instances.Data.TotalRam))
	res.WriteString("\n")
	res.WriteString("Total Disk: ")
	res.WriteString(strconv.Itoa(instances.Data.TotalDisk))
	res.WriteString("\n---\n")
	for _, item := range instances.Data.Instances {
		res.WriteString("Hostname: ")
		res.WriteString(item.Hostname)
		res.WriteString("\n")
		res.WriteString("Ips: ")
		res.WriteString(item.Ips)
		res.WriteString("\n")

		instance, err := cn.GetInstance(item.Id)
		if err != nil {
			return []result.Result{result.ErrorResult(err)}
		}

		data, _ := json.Marshal(instance.Data.Instances.Billing.Monthly)
		res.WriteString("Billing: ")
		res.Write(data)
		res.WriteString("\n\n")
	}

	return []result.Result{result.MessageResult(res.String())}
}
