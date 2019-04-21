package combination

import (
	"time"
)

type Combination struct {
	cache   map[int]map[int]uint64
	updated time.Time // this is used to check for improver caching.
}

func NewCombination() *Combination {
	c := make(map[int]map[int]uint64)

	return &Combination{
		cache: c,
	}
}

func (c *Combination) Choose(n, k int) uint64 {
	// base case
	if k == 0 {
		return 1
	}

	// first check the cache
	nmap, ok := c.cache[n]

	// if map for key n does not exist, create one, compute nCk, add it as value for key k, and return nCk
	if !ok {
		newNmap := make(map[int]uint64)
		value := c.Choose(n-1, k-1) * uint64(n) / uint64(k)
		newNmap[k] = value
		c.cache[n] = newNmap
		c.updated = time.Now()
		return value
	}

	// map exists for key n.  Check map for value of key k
	nck, ok := nmap[k]

	// if there is no value for key k, compute nCk, add it as value for key k, and return nCk
	if !ok {
		value := c.Choose(n-1, k-1) * uint64(n) / uint64(k)
		nmap[k] = value
		c.updated = time.Now()
		return value
	}

	// the combination is in the cache
	return nck
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
