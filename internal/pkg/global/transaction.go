package global

import (
	"github.com/dtm-labs/dtm/dtmgrpc"
	"github.com/tsundata/assistant/internal/pkg/config"
)

type Transaction struct {
	dtm string
}

func NewTransaction(conf *config.AppConfig) *Transaction {
	return &Transaction{
		dtm: conf.SvcAddr.Dtm,
	}
}

func (t *Transaction) TCC(tccFunc dtmgrpc.TccGlobalFunc) (string, error) {
	gid := dtmgrpc.MustGenGid(t.dtm)
	return gid, dtmgrpc.TccGlobalTransaction(t.dtm, gid, tccFunc)
}
