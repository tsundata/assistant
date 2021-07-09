package agent

import (
	"reflect"
	"testing"

	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func TestFetchGithubStarred(t *testing.T) {
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
			if got := FetchGithubStarred(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchGithubStarred() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchGithubStargazers(t *testing.T) {
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
			if got := FetchGithubStargazers(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchGithubStarred() = %v, want %v", got, tt.want)
			}
		})
	}
}
