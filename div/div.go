// Package div provides a function for converting integer division
// into integer multiplication for chosen denominator values. This
// can be generalized to other operations based on the equation:
//
// 	a = bq + r
//
package div

// Invert returns three coefficients to compute quotients on divisor b using the
// following arithmetic operation: q = (a*B+r) >> k. The dividend, a, is an integer
// in the domain [1, 2**(N-1)).
func InvertN(b, N uint64) (B, r, k uint64) {
	const _1 = uint64(1)
	e1 := _1 << (N - 1)
	k = N + log2(b)
	s := (_1 << k)
	B = (s / b)
	return B, (e1 / b) * (s - B*b), k
}

// Invert returns three coefficients to compute quotients on divisor b using the
// following arithmetic operation: q = (a*B+r) >> k. The dividend, a, is an integer
// in the domain [1, 2**31).
func Invert32(b uint64) (B, r, k uint64) {
	return InvertN(b, 32)
}

// avoids dependency on math/bits for older go releases
func log2(n uint64) (e uint64) {
	if n >= 1<<32 {
		n >>= 32
		e = 32
	}
	if n >= 1<<16 {
		n >>= 16
		e += 16
	}
	if n >= 1<<8 {
		n >>= 8
		e += 8
	}
	return e + len8tab[n] - 1
}

var len8tab = [256]uint64{
	0, 1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
}

// mod is an example function; really shouldn't be using
// it since the function call overhead defeats the purpose
func mod(a, b uint64) uint64 {
	return a - b*div(a, b)
}

// div is an example function; really shouldn't be using
// it since the function call overhead defeats the purpose
func div(a, b uint64) uint64 {
	B, r, k := Invert32(b)
	return (a*B + r) >> k
}
