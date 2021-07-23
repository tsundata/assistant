package config

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAppConfig_Config(t *testing.T) {
	c, err := CreateAppConfig("test")
	if err != nil {
		t.Fatal(err)
	}
	res, err := c.GetConfig(context.Background(), "common")
	if err != nil {
		t.Fatal(err)
	}
	require.NotEqual(t, "", res)
}

func TestAppConfig_Setting(t *testing.T) {
	c, err := CreateAppConfig("test")
	if err != nil {
		t.Fatal(err)
	}
	err = c.SetSetting(context.Background(), "key", "value")
	if err != nil {
		t.Fatal(err)
	}
	res, err := c.GetSetting(context.Background(), "key")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "value", res)

	err = c.SetSetting(context.Background(), "key2", "value2")
	if err != nil {
		t.Fatal(err)
	}
	res2, err := c.GetSettings(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, len(res2) >= 2)
}
