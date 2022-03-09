package util

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestIsApartDay(t *testing.T) {
	type args struct {
		in0 string
		in1 string
		in2 string
		in3 int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"case1",
			args{"2021-06-22 15:35", "", "2021-06-22 15:35", 1},
			false,
			false,
		},
		{
			"case2",
			args{"2021-06-22 15:35", "", "2021-06-22 15:35", 5},
			false,
			false,
		},
		{
			"case3",
			args{"2021-06-22 15:35", "2021-06-22 15:35", "2021-06-23 15:50", 1},
			true,
			false,
		},
		{
			"case4",
			args{"2021-06-22 15:35", "2021-06-22 15:35", "2021-06-27 15:50", 5},
			true,
			false,
		},
		{
			"case5",
			args{"2021-06-22 15:35", "2021-06-22 15:35", "2021-06-23 15:50", 5},
			false,
			false,
		},
		{
			"case6",
			args{"2021-12-22 15:35", "2022-12-22 15:35", "2022-02-10 15:50", 30},
			false,
			false,
		},
		{
			"case7",
			args{"2021-12-27 15:35", "", "2021-12-22 15:35", 30},
			false,
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsApartDay(tt.args.in0, tt.args.in1, tt.args.in2, tt.args.in3)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDaily error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsDaily() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDaily(t *testing.T) {
	type args struct {
		in0 string
		in1 string
		in2 string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"case1",
			args{"2021-06-22 15:35", "", "2021-06-22 15:35"},
			false,
			false,
		},
		{
			"case2",
			args{"2021-06-22 15:35", "2021-06-22 15:35", "2021-06-22 15:50"},
			false,
			false,
		},
		{
			"case3",
			args{"2021-06-22 15:35", "2021-06-22 15:35", "2021-06-23 15:50"},
			true,
			false,
		},
		{
			"case4",
			args{"2021-06-29 17:57", "2021-06-30 18:47", "2021-06-30 18:48"},
			false,
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsDaily(tt.args.in0, tt.args.in1, tt.args.in2)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDaily error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsDaily() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWeekly(t *testing.T) {
	type args struct {
		in0 string
		in1 string
		in2 string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"case1",
			args{"2021-06-22 15:35", "", "2021-06-22 15:35"},
			false,
			false,
		},
		{
			"case2",
			args{"2021-06-22 15:35", "2021-06-22 15:35", "2021-06-22 15:50"},
			false,
			false,
		},
		{
			"case3",
			args{"2021-06-22 15:35", "2021-06-22 15:35", "2021-06-29 15:50"},
			true,
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsWeekly(tt.args.in0, tt.args.in1, tt.args.in2)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDaily error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsDaily() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMonthly(t *testing.T) {
	type args struct {
		in0 string
		in1 string
		in2 string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"case1",
			args{"2021-06-22 15:35", "", "2021-06-22 15:35"},
			false,
			false,
		},
		{
			"case2",
			args{"2021-06-22 15:35", "2021-06-22 15:35", "2021-07-12 15:50"},
			false,
			false,
		},
		{
			"case3",
			args{"2021-06-22 15:35", "2021-06-22 15:35", "2021-07-22 15:50"},
			true,
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsMonthly(tt.args.in0, tt.args.in1, tt.args.in2)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDaily error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsDaily() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAnnually(t *testing.T) {
	type args struct {
		in0 string
		in1 string
		in2 string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"case1",
			args{"2021-06-22 15:35", "", "2021-06-22 15:35"},
			false,
			false,
		},
		{
			"case2",
			args{"2021-06-22 15:35", "2021-06-22 15:35", "2022-01-12 15:50"},
			false,
			false,
		},
		{
			"case3",
			args{"2021-06-22 15:35", "2021-06-22 15:35", "2022-06-22 15:50"},
			true,
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsAnnually(tt.args.in0, tt.args.in1, tt.args.in2)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDaily error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsDaily() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNow(t *testing.T) {
	require.Regexp(t, `2\d{3}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, Now())
}

func TestFormat(t *testing.T) {
	require.Equal(t, "2022-03-09 11:15:48", Format(1646795748))
}
