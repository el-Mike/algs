package hashtable

import (
	"fmt"
	"sync"

	"github.com/el-Mike/algs/data-structures/list"
)

type HashtableItem struct {
	Key   int
	Value int
}

func NewHashtableItem(key, value int) *HashtableItem {
	return &HashtableItem{
		Key:   key,
		Value: value,
	}
}

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

// hash - returns key hash based on previously set size and hashStrategy.
func (h *ChainingHashtable) hash(key int) int {
	hasher := GetHasher(h.hashStrategy)

	return hasher(key, h.size)
}

// Insert - adds an element to the Hashmap.
func (h *ChainingHashtable) Insert(key, value int) error {
	index := h.hash(key)

	item := NewHashtableItem(key, value)
	node := list.NewNode(item, nil)

	if h.table[index] == nil {
		h.table[index] = list.NewLinkedList(node)
	} else {
		if h.table[index].Find(h.findCallbackFactory(key)) != nil {
			return NewKeyAlreadyExistsError()
		}
		h.table[index].Prepend(item)
	}

	return nil
}

// Get - returns an item for given key. It searches underlying LinkedList
// to get the actual item.
func (h *ChainingHashtable) Get(key int) int {
	index := h.hash(key)

	if h.table[index] == nil {
		return 0
	}

	node := h.table[index].Find(h.findCallbackFactory(key))

	if node == nil {
		return 0
	}

	item := h.castHashtableItem(node.Data)

	if item == nil {
		return 0
	}

	return item.Value
}

// Remove - removes an item under given key.
func (h *ChainingHashtable) Remove(key int) error {
	index := h.hash(key)

	if h.table[index] == nil {
		return NewKeyDoesNotExistError()
	}

	// If the element to be removed is the head, we set the head to the
	// next list element (or nil if empty).
	if h.castHashtableItem(h.table[index].Head.Data).Key == key {
		h.table[index].Head = h.table[index].Head.Next

		return nil
	}

	found := false
	h.table[index].ForEach(func(node *list.Node) {
		// When we find preceding Node, we set it's Next to succeeding Node.
		if h.castHashtableItem(node.Next.Data).Key == key {
			node.Next = node.Next.Next
			found = true
		}
	})

	if !found {
		return NewKeyDoesNotExistError()
	}

	return nil
}

// Print - prints Hashtable's LinkedLists as an arrays.
func (h *ChainingHashtable) Print(printer list.NodePrinter) {
	for _, l := range h.table {
		if l != nil {
			l.Print(printer)
		}
	}
}

// findCallbackFactory - helper function for find Callback based on item's key.
func (h *ChainingHashtable) findCallbackFactory(key int) list.FindCallback {
	return func(node *list.Node) bool {
		return h.castHashtableItem(node.Data).Key == key
	}
}

func (h *ChainingHashtable) castHashtableItem(data interface{}) *HashtableItem {
	return data.(*HashtableItem)
}

func RunChainingHashtable() {
	hashtable := NewHashtable(10, MOD_HASHING)

	testData := GetSmallTestDataCopy()

	for i, x := range testData {
		hashtable.Insert((x + i + 10), x)
	}

	hashtable.Print(printNode)
}

func printNode(node *list.Node) {
	fmt.Printf("%v, ", node.Data.(*HashtableItem).Value)
}
