package hashtable

import (
	"sync"

	"github.com/el-Mike/algs/data-structures/list"
)

type ChainingHashtable struct {
	// Makes hashtable implementation thread-safe.
	sync.RWMutex

	size         int
	hashStrategy HashStrategy
	table        []*list.LinkedList
}

// NewHashtable - returns new Hashtable instance.
func NewHashtable(size int, hashStrategy HashStrategy) *ChainingHashtable {
	return &ChainingHashtable{
		size:         size,
		hashStrategy: hashStrategy,
		table:        make([]*list.LinkedList, size),
	}
}

// Insert - adds an element to the Hashmap.
func (h *ChainingHashtable) Insert(key, value int) error {
	hasher := GetHasher(h.hashStrategy)
	index := hasher(key, h.size)

	item := list.NewNode(value, nil)

	if h.table[index] == nil {
		h.table[index] = list.NewLinkedList(item)
	} else {
		h.table[index].Append(item)
	}

	return nil
}
