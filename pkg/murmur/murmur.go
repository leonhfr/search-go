// Package murmur implements the Murmur3 hash function based on
// http://en.wikipedia.org/wiki/MurmurHash
package murmur

const (
	c1 = 0xcc9e2d51
	c2 = 0x1b873593
	c3 = 0x85ebca6b
	c4 = 0xc2b2ae35
	nh = 0xe6546b64
)

// Sum returns a hash from the provided key using the seed.
func Sum(key string, seed uint32) (hash uint32) {
	hash = seed

	nbytes := len(key) / 4 * 4
	for i := 0; i < nbytes; i += 4 {
		k := uint32(key[i]) | uint32(key[i+1])<<8 |
			uint32(key[i+2])<<16 | uint32(key[i+3])<<24

		k *= c1
		k = (k << 15) | (k >> 17)
		k *= c2
		hash ^= k

		hash = (hash << 13) | (hash >> 19)
		hash = hash*5 + nh
	}

	var remaining uint32
	switch len(key) & 3 {
	case 3:
		remaining += uint32(key[nbytes+2]) << 16
		fallthrough
	case 2:
		remaining += uint32(key[nbytes+1]) << 8
		fallthrough
	case 1:
		remaining += uint32(key[nbytes])
		remaining *= c1

		remaining = (remaining << 15) | (remaining >> 17)
		remaining = remaining * c2
		hash ^= remaining
	}

	hash ^= uint32(len(key))
	hash ^= hash >> 16
	hash *= c3
	hash ^= hash >> 13
	hash *= c4
	hash ^= hash >> 16

	return
}
