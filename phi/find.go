package phi

import "math/big"

var (
	one    = big.NewInt(1)
	two    = big.NewInt(2)
	negtwo = big.NewInt(-2)
)

// Find directly computes φ(N)
func Find(N *big.Int) *big.Int {
	g := one
	E := guess(N, g, new(big.Int))
	E0 := new(big.Int).Exp(two, E, N)
	E0 = E0.ModInverse(E0, N)
	L := log2(N)
	i := int64(0)
	for !smooth2(E0) {
		E0.Lsh(E0, L)
		E0.Mod(E0, N)
		i++
	}
	E = E.Neg(E)
	E = E.Sub(E, big.NewInt(i*int64(L)-int64(log2(E0))))
	return E.Neg(E)
}

func log2(N *big.Int) uint {
	return uint(N.BitLen() - 1)
}

func guess(N, g, E *big.Int) *big.Int {
	// n - sqrt(n)*2 + 1
	return E.Add(N, E.Sub(E.Mul(E.Sqrt(N), negtwo), one))
}

// smooth2 returns true iff n is 2-smooth
func smooth2(n *big.Int) bool {
	b := n.Bytes()
	if b[0] != 1 {
		return false
	}
	for i := 1; i < len(b); i++ {
		if b[i] != 0 {
			return false
		}
	}
	return true
}
