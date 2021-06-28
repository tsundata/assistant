package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAppConfig_Watch(t *testing.T) {
	t.Skip() // todo
	c, err := CreateAppConfig("test")
	if err != nil {
		t.Fatal(err)
	}
	c.Watch()
}

func TestAppConfig_Config(t *testing.T) {
	c, err := CreateAppConfig("test")
	if err != nil {
		t.Fatal(err)
	}
	res, err := c.GetConfig("common")
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
	err = c.SetSetting("key", "value")
	if err != nil {
		t.Fatal(err)
	}
	res, err := c.GetSetting("key")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "value", res)

	err = c.SetSetting("key2", "value2")
	if err != nil {
		t.Fatal(err)
	}
	res2, err := c.GetSettings()
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, len(res2) >= 2)
}
