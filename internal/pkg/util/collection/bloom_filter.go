package collection

import (
	"github.com/go-redis/redis/v8"
	"github.com/spaolacci/murmur3"
	"sync"
)

type BloomFilter struct {
	bitMap    *BinSet
	mu        sync.RWMutex
	size      uint
	hashCount uint
}

func NewBloomFilter(rdb *redis.Client, key string, size uint, hashCount uint) *BloomFilter {
	bf := BloomFilter{bitMap: NewBinSet(rdb, key), size: size, hashCount: hashCount, mu: sync.RWMutex{}}
	return &bf
}

func (bf *BloomFilter) Lookup(value string) bool {
	bf.mu.RLock()
	defer bf.mu.RUnlock()
	for i := uint(0); i < bf.hashCount; i++ {
		result := bf.hash(value) % uint64(bf.size)
		if !bf.bitMap.Test(int64(result)) {
			return false
		}
	}

	return true
}

func (bf *BloomFilter) Add(value string) {
	bf.mu.Lock()
	defer bf.mu.Unlock()
	for i := uint(0); i < bf.hashCount; i++ {
		result := bf.hash(value) % uint64(bf.size)
		bf.bitMap.SetTo(int64(result), 1)
	}
}

func (bf *BloomFilter) hash(value string) uint64 {
	m := murmur3.New64()
	_, _ = m.Write([]byte(value))
	return m.Sum64()
}
