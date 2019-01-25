package prob

import (
	"math/bits"
	"math/rand"
)

type Coin struct {
	ctr   int
	state uint64
	rand.Rand
}

// Flip returns the outcome of a random event
// with probability n/m
func (c *Coin) Flip(n, m int) (heads bool) {
	if c.ctr < m {
		c.state = uint64(c.Rand.Int63()) << 1
		c.ctr = 63
	}
	mask := ^((^uint64(0)) >> uint(m+1))
	heads = bits.OnesCount64(c.state&mask) > n
	c.state <<= uint(m)
	c.ctr -= m
	return heads
}
