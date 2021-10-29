package queue

import "sync"

type PriorityQueue struct {
	// Makes queue implementation thread-safe.
	sync.RWMutex
}
