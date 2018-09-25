package div

// Egcd outputs the greatest common divisor of a and b, as
// well as the Bezout coefficients, x and y. The coefficients
// are components of g, as per: ax + by = gcd(a,b) == g
func Egcd(a, b int) (g, x, y int) {
	y = 1
	u, v := 1, 0
	for a != 0 {
		q, r := b/a, b%a
		m, n := x-u*q, y-v*q
		b, a, x, y, u, v = a, r, u, v, m, n
	}
	return b, x, y
}
