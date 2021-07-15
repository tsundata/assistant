package rqlite

import (
	"github.com/google/wire"
	"github.com/rqlite/gorqlite"
	"github.com/tsundata/assistant/internal/pkg/config"
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
