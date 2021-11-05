package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainingHashtable(t *testing.T) {
	hashtable := NewHashtable(10, MOD_HASHING)

	hashtable.Insert(12, 1)
	hashtable.Insert(13, 2)

	assert.Equal(t, 1, hashtable.Get(12))
	assert.Equal(t, 2, hashtable.Get(13))

	hashtable.Insert(22, 3)
	hashtable.Insert(32, 4)

	assert.Equal(t, 3, hashtable.Get(22))
	assert.Equal(t, 4, hashtable.Get(32))

	assert.NotEqual(t, 3, hashtable.Get(12))

	hashtable.Remove(12)

	assert.Equal(t, 0, hashtable.Get(12))
}
