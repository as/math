package div

import "testing"

const lim = 1 << 12

func TestInverseRoots(t *testing.T) {
	var tab = [65536]struct {
		a uint64
		d uint64
		m uint64
	}{}
	for n := range tab {
		{
			n := uint64(n)
			if n != 0 {
				tab[n].a, tab[n].d, tab[n].m = Invert32(n)
			}
		}
	}

	// anything higher takes a noticable amount of time
	t.Run("Mod", func(t *testing.T) {
		right, total := 0, 0
		for a := uint64(1); a < lim; a++ {
			for b := uint64(1); b < lim; b++ {
				total++
				if have, want := mod(a, b), a%b; have == want {
					right++
				} else {
					t.Fatalf("bad modulus: %d/%d: %d,%d ", a, b, have, want)
				}
			}
		}
		t.Logf("%d divisions (%d correct)", total, right)
	})

	// anything higher takes a noticable amount of time
	t.Run("Slow", func(t *testing.T) {
		right, total := 0, 0
		for a := uint64(1); a < lim; a++ {
			for b := uint64(1); b < lim; b++ {
				total++
				if have, want := div(a, b), a/b; have == want {
					right++
				} else {
					t.Fatalf("bad division: %d/%d: %d,%d acc(%0.3f)\n", a, b, have, want, float64(have)/float64(want))
				}
			}
		}
		t.Logf("%d divisions (%d correct)", total, right)
	})

	tabdiv := func(a, b uint64) uint64 {
		t := tab[b]
		return (t.a*a + t.d) >> t.m
	}
	t.Run("Lut", func(t *testing.T) {
		right, total := 0, 0
		for a := uint64(1); a < lim; a++ {
			for b := uint64(1); b < lim; b++ {
				total++
				if have, want := tabdiv(a, b), a/b; have == want {
					right++
				} else {
					t.Fatalf("bad division: %d/%d: %d,%d acc(%0.3f)\n", a, b, have, want, float64(have)/float64(want))
				}
			}
		}
		t.Logf("%d divisions (%d correct)", total, right)
	})
}
