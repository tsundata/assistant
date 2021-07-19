package rqlite

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/gorqlite"
)

type Conn struct {
	nr   *newrelic.App
	conn gorqlite.Connection
}

func New(c *config.AppConfig, nr *newrelic.App) (*Conn, error) {
	conn, err := gorqlite.Open(c.Rqlite.Url)
	if err != nil {
		return nil, err
	}
	err = conn.SetConsistencyLevel("strong")
	if err != nil {
		return nil, err
	}
	return &Conn{nr: nr, conn: conn}, nil
}

func (c *Conn) Query(sqlStatements []string) ([]gorqlite.QueryResult, error) {
	nxt := c.nr.StartTransaction("rqlite/query")
	defer nxt.End()
	for _, statement := range sqlStatements {
		segment := nxt.StartSegment(statement)
		segment.End()
	}

	return c.conn.Query(sqlStatements)
}

func (c *Conn) QueryOne(sqlStatement string) (gorqlite.QueryResult, error) {
	nxt := c.nr.StartTransaction("rqlite/query")
	defer nxt.End()
	segment := nxt.StartSegment(sqlStatement)
	defer segment.End()

	return c.conn.QueryOne(sqlStatement)
}

func (c *Conn) Write(sqlStatements []string) ([]gorqlite.WriteResult, error) {
	nxt := c.nr.StartTransaction("rqlite/write")
	defer nxt.End()
	for _, statement := range sqlStatements {
		segment := nxt.StartSegment(statement)
		segment.End()
	}

	return c.conn.Write(sqlStatements)
}

func (c *Conn) WriteOne(sqlStatement string) (gorqlite.WriteResult, error) {
	nxt := c.nr.StartTransaction("rqlite/write")
	defer nxt.End()
	segment := nxt.StartSegment(sqlStatement)
	defer segment.End()

	return c.conn.WriteOne(sqlStatement)
}

var ProviderSet = wire.NewSet(New)
