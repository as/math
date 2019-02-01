package mod

import "math/bits"

// montgomery computes a mod b via the montgomery
// reduction. The arguments must satisfy the inequality
// 0 <= a <= b**2 and b must be an odd integer.
// runtime properties: 2k**2+1 shifts, k**2 adds
//
// this function returns a mod b multiplied by a constant
// k, where k is th enumber of bits
func montgomery(a, b int) int {
	for n := bits.Len(uint(a)); n > 0; n-- {
		if a&1 == 1 {
			a += b
		}
		a >>= 1
	}
	return a
}

// montgomery2 is an optimized version of the montgomery
// function above. The inequality must hold for a and b:
// 2**(width of a) > b
func montgomery2(a, b int) int {
	n := bits.Len(uint(a))
	for k := 1; k < n; k++ {
		if a>>uint(k-1)&1 == 1 {
			a += 1 << uint(b)
		}
	}
	return a >> uint(n)
}
