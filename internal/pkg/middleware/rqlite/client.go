package rqlite

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/gorqlite"
)

func New(c *config.AppConfig) (gorqlite.Connection, error) {
	conn, err := gorqlite.Open(c.Rqlite.Url)
	if err != nil {
		return gorqlite.Connection{}, err
	}
	err = conn.SetConsistencyLevel("strong")
	if err != nil {
		return gorqlite.Connection{}, err
	}
	return conn, nil
}

var ProviderSet = wire.NewSet(New)
