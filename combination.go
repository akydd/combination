package combination

import (
	"math/big"
	"time"
)

type Combination struct {
	cache   map[int]map[int]*big.Int
	updated time.Time // this is used to check for improver caching.
}

func NewCombination() *Combination {
	c := make(map[int]map[int]*big.Int)
	//fmt.Printf("Max value is %d\n", uint64(math.MaxUint64))

	return &Combination{
		cache: c,
	}
}

func (c *Combination) Choose(n, k int) big.Int {
	// base case
	if k == 0 {
		return *big.NewInt(1)
	}

	// first check the cache
	nmap, ok := c.cache[n]

	// if map for key n does not exist, create one, compute nCk, add it as value for key k, and return nCk
	if !ok {
		newNmap := make(map[int]*big.Int)
		// Choose(n, k) = Choose(n-1, k-1) * n / k
		var value big.Int
		p := c.Choose(n-1, k-1)
		value.Div(value.Mul(&p, big.NewInt(int64(n))), big.NewInt(int64(k)))

		newNmap[k] = &value
		c.cache[n] = newNmap
		c.updated = time.Now()
		return value
	}

	// map exists for key n.  Check map for value of key k
	nck, ok := nmap[k]

	// if there is no value for key k, compute nCk, add it as value for key k, and return nCk
	if !ok {
		var value big.Int
		p := c.Choose(n-1, k-1)
		value.Div(value.Mul(&p, big.NewInt(int64(n))), big.NewInt(int64(k)))

		nmap[k] = &value
		c.updated = time.Now()
		return value
	}

	// the combination is in the cache
	return *nck
}

func (c *Combination) IsNcached(n int) bool {
	_, ok := c.cache[n]
	return ok
}

func (c *Combination) IsCached(n, k int) bool {
	nCache, ok := c.cache[n]
	if !ok {
		return false
	}

	_, ok = nCache[k]
	return ok
}

func (c *Combination) GetUpdated() time.Time {
	return c.updated
}
