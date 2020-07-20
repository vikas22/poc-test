package utils

import "sync/atomic"

type UID uint64

func (c *UID) Inc() uint64 {
	return atomic.AddUint64((*uint64)(c), 1)
}

func (c *UID) Get() uint64 {
	return atomic.LoadUint64((*uint64)(c))
}
