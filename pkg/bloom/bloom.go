package bloom

import (
	"crypto/rand"
	"math"
	"math/big"

	"github.com/leonhfr/search-go/pkg/murmur"
)

const (
	p = 1e-3
	k = 14
)

type Filter struct {
	Bitset []uint64
	Seeds  [k]uint32
}

func New(n int) Filter {
	bits := optimalBits(n)
	size := bits / 64
	if bits%64 > 0 {
		size++
	}

	return Filter{
		Bitset: make([]uint64, size),
		Seeds:  generateSeeds(),
	}
}

func (f *Filter) Add(value string) {
	bits := 64 * len(f.Bitset)

	for _, seed := range f.Seeds {
		hash := murmur.Sum(value, seed)
		bit := int(hash) % bits
		f.Bitset[bit/64] |= 1 << (bit % 64)
	}
}

func (f *Filter) Query(query string) bool {
	bits := 64 * len(f.Bitset)

	for _, seed := range f.Seeds {
		hash := murmur.Sum(query, seed)
		bit := int(hash) % bits
		set := f.Bitset[bit/64] & (1 << (bit % 64))
		if set == 0 {
			return false
		}
	}

	return true
}

func generateSeeds() [k]uint32 {
	var s [k]uint32
	for i := 0; i < k; i++ {
		v, _ := rand.Int(rand.Reader, big.NewInt(math.MaxUint32))
		s[i] = uint32(v.Uint64())
	}
	return s
}

func optimalBits(n int) int {
	m := -float64(n) * math.Log(p) / math.Pow(math.Ln2, 2)
	return int(math.Ceil(m))
}
