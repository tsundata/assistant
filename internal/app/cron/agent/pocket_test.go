package agent

import (
	"reflect"
	"testing"

	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func TestFetchPocket(t *testing.T) {
	type args struct {
		ctx rulebot.IContext
	}
	tests := []struct {
		name string
		args args
		want []result.Result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchPocket(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchPocket() = %v, want %v", got, tt.want)
			}
		})
	}
}
