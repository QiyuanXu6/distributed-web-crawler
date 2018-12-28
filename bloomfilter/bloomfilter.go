package bloomfilter

import (
	"crypto/sha256"
	"github.com/gomodule/redigo/redis"
	"github.com/spaolacci/murmur3"
)

// Bloomfilter interface
type Bloomfilter interface {
	Add([]byte)
	Check([]byte) bool
}

// Redis Bloomfilter implementation
type RedisBF struct {
	client redis.Conn
	n      uint // len of buckets
	k      uint // number of hash func
	key    string
}

func (bf *RedisBF) Add(data []byte) {
	for i := uint(0); i < bf.k; i++ {
		val := uint(hash(data, uint32(i))) % bf.n
		bf.client.Do("SETBIT", bf.key, val, 1)
	}
}

func (bf *RedisBF) Check(data []byte) bool {
	for i := uint(0); i < bf.k; i++ {
		val := uint(hash(data, uint32(i))) % bf.n
		reply, _ := redis.Int(bf.client.Do("GETBIT", bf.key, val))
		if reply == 0 {
			return false
		}
	}
	return true
}

func NewRedisBF(client redis.Conn, n, k uint, key string) *RedisBF {
	bf := &RedisBF{
		client: client,
		n:      n,
		k:      k,
		key:    key,
	}

	client.Do("DEL", bf.key)
	return bf
}

func hash(data []byte, seed uint32) uint64 {
	// length to same
	sha := sha256.Sum256(data)
	m := murmur3.New64WithSeed(seed)
	m.Write(sha[:])
	return m.Sum64()
}
