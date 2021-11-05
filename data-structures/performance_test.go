package datastructures

import (
	"testing"

	"github.com/el-Mike/algs/data-structures/hashtable"
	"github.com/el-Mike/algs/data-structures/list"
	testdata "github.com/el-Mike/algs/data-structures/test_data"
)

func BenchmarkChainingHashtableInsert(b *testing.B) {
	testData := testdata.GetTestDataCopy()
	hash := hashtable.NewHashtable(500, hashtable.MOD_HASHING)

	for i, x := range testData {
		hash.Insert((x + i + 3), x)
	}

	for i := 0; i < b.N; i++ {
		if err := hash.Insert((i + 3), i); err != nil {
			continue
		}
	}
}

func BenchmarkLinkedListInsert(b *testing.B) {
	testData := testdata.GetTestDataCopy()
	l := list.NewLinkedList(nil)

	for _, x := range testData {
		l.Append(x)
	}

	for i := 0; i < b.N; i++ {
		l.Append(i)
	}
}

func BenchmarkChainingHashtableGet(b *testing.B) {
	testData := testdata.GetTestDataCopy()
	hash := hashtable.NewHashtable(500, hashtable.MOD_HASHING)

	for i, x := range testData {
		hash.Insert((x + i + 3), x)
	}

	for i := 0; i < b.N; i++ {
		hash.Get((i))
	}
}

func BenchmarkLinkedListGet(b *testing.B) {
	testData := testdata.GetTestDataCopy()
	l := list.NewLinkedList(nil)

	for _, x := range testData {
		l.Append(x)
	}

	for i := 0; i < b.N; i++ {
		l.Find(func(node *list.Node) bool {
			return node.Data.(int) == i
		})
	}
}
