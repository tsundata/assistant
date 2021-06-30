package service

import (
	"reflect"
	"testing"

	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func TestCreateInitServerFn(t *testing.T) {
	type args struct {
		ps *Middle
	}
	tests := []struct {
		name string
		args args
		want rpc.InitServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateInitServerFn(tt.args.ps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInitServerFn() = %v, want %v", got, tt.want)
			}
		})
	}
}
