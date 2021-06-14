package collection

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type BinSet struct {
	rdb *redis.Client
	key string
}

func NewBinSet(rdb *redis.Client, key string) *BinSet {
	return &BinSet{rdb: rdb, key: key}
}

func (b *BinSet) SetTo(set int64, state int) {
	b.rdb.SetBit(context.Background(), b.key, set, state)
}

func (b *BinSet) Test(set int64) bool {
	i := b.rdb.GetBit(context.Background(), b.key, set)
	return i.Val() == 1
}
