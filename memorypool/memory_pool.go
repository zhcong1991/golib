package memorypool

import (
	"sync"
)

type MemoryPool struct {
	pool      sync.Pool
	blockSize int
}

func (p *MemoryPool) Put(b []byte) {
	p.pool.Put(b)
}

func (p *MemoryPool) Get() []byte {
	return p.pool.Get().([]byte)
}

func NewMemoryPool(blockSize int) *MemoryPool {
	p := &MemoryPool{
		blockSize: blockSize,
	}
	p.pool.New = func() any {
		return make([]byte, p.blockSize)
	}
	return p
}
