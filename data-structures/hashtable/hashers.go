package hashtable

type HashStrategy string
type Hasher func(key, size int) int

const (
	MOD_HASHING            HashStrategy = "Mod"
	MULTIPLICATION_HASHING HashStrategy = "Multiplication"
)

func GetHasher(strategy HashStrategy) Hasher {
	switch strategy {
	case MOD_HASHING:
		return ModHash
	case MULTIPLICATION_HASHING:
		return nil
	default:
		return nil
	}
}

// ModHash - uses simple modulo operation on input key to get hashed index.
func ModHash(key, size int) int {
	return key % size
}
