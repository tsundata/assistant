package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/tsundata/assistant/api/pb"
)

func TestFinance_CreateBill(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *pb.BillRequest
	}
	tests := []struct {
		name    string
		f       *Finance
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.CreateBill(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Finance.CreateBill() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Finance.CreateBill() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFinance_GetBill(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *pb.BillRequest
	}
	tests := []struct {
		name    string
		f       *Finance
		args    args
		want    *pb.BillReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.GetBill(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Finance.GetBill() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Finance.GetBill() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFinance_GetBills(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *pb.BillRequest
	}
	tests := []struct {
		name    string
		f       *Finance
		args    args
		want    *pb.BillsReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.GetBills(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Finance.GetBills() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Finance.GetBills() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFinance_DeleteBill(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *pb.BillRequest
	}
	tests := []struct {
		name    string
		f       *Finance
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.DeleteBill(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Finance.DeleteBill() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Finance.DeleteBill() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFinance_GetFund(t *testing.T) {
	t.SkipNow()
	s := NewFinance()

	type args struct {
		ctx     context.Context
		payload *pb.TextRequest
	}
	tests := []struct {
		name    string
		f       *Finance
		args    args
		wantErr bool
	}{
		{"case1", s, args{context.Background(), &pb.TextRequest{Text: "000001"}}, false},
		{"case2", s, args{context.Background(), &pb.TextRequest{Text: "003171"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.f.GetFund(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Finance.GetFund() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestFinance_GetStock(t *testing.T) {
	t.SkipNow()
	s := NewFinance()

	type args struct {
		ctx     context.Context
		payload *pb.TextRequest
	}
	tests := []struct {
		name    string
		f       *Finance
		args    args
		wantErr bool
	}{
		{"case1", s, args{context.Background(), &pb.TextRequest{Text: "sz000001"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.f.GetStock(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Finance.GetStock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
