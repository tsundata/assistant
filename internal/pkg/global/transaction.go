package global

import (
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/yedf/dtm/dtmgrpc"
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
