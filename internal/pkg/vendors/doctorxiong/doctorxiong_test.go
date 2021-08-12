package doctorxiong

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestDoctorxiong_GetFundDetail(t *testing.T) {
	dx := NewDoctorxiong("")
	resp, err := dx.GetFundDetail("000001", "2021-08-01", "2021-08-31")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, http.StatusOK, resp.Code)
}

func TestDoctorxiong_GetFundDetail2(t *testing.T) {
	dx := NewDoctorxiong("")
	resp, err := dx.GetFundDetail("003171", "2021-08-01", "2021-08-31")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, http.StatusOK, resp.Code)
}

func TestDoctorxiong_GetStockDetail(t *testing.T) {
	dx := NewDoctorxiong("")
	resp, err := dx.GetStockDetail("sz000001")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, http.StatusOK, resp.Code)
}
