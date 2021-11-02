package hashtable

import "sync"

type HashStrategy string
type Hasher func(key int) int

type HashItem struct {
	key  int
	data int
}

type Hashtable struct {
	// Makes hashtable implementation thread-safe.
	sync.RWMutex

	size         int
	hashStrategy HashStrategy
	table        []*HashItem
}

const (
	MOD_HASHING            HashStrategy = "Mod"
	MULTIPLICATION_HASHING HashStrategy = "Multiplication"
)

// NewHashItem - returns new Hashitem instance.
func NewHashItem(key, data int) *HashItem {
	return &HashItem{
		key:  key,
		data: data,
	}
}

// NewHashtable - returns new Hashtable instance.
func NewHashtable(size int, hashStrategy HashStrategy) *Hashtable {
	return &Hashtable{
		size:         size,
		hashStrategy: hashStrategy,
		table:        make([]*HashItem, size),
	}
}

// Insert - adds an element to the Hashmap.
func (h *Hashtable) Insert(key, value int) error {
	hasher := h.GetHasher()

	index := hasher(key)

	if h.table[index] != nil && h.table[index].key == key {
		return NewKeyNotEmptyError()
	}

	item := NewHashItem(key, value)
	h.table[index] = item

	return nil
}

func (h *Hashtable) GetHasher() Hasher {
	switch h.hashStrategy {
	case MOD_HASHING:
		return h.modHash
	case MULTIPLICATION_HASHING:
		return nil
	default:
		return nil
	}
}

// ModHash - uses simple modulo operation on input key to get hashed index.
func (h *Hashtable) modHash(key int) int {
	return key % h.size
}
