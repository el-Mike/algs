package hashtable

type HashStrategy string
type Hasher func(key, size int) int

const (
	MOD_HASHING                  HashStrategy = "Mod"
	KNUTH_MULTIPLICATION_HASHING HashStrategy = "KnuthMultiplication"
)

func GetHasher(strategy HashStrategy) Hasher {
	switch strategy {
	case MOD_HASHING:
		return ModHash
	case KNUTH_MULTIPLICATION_HASHING:
		return KnuthMultiplicationHash
	default:
		return nil
	}
}

// ModHash - uses simple modulo operation on input key to get hashed index.
func ModHash(key, size int) int {
	return key % size
}

func KnuthMultiplicationHash(key, size int) int {
	// From Knuth's "Sorting and Searching"
	// alpha := math.Pow(2, 32) * (math.Sqrt(5)-1) / 2
	alpha := 2654435769

	return ((key * alpha) >> 16) % size
}
