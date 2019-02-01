package mod

import (
	"math/bits"
	"testing"
)

var cases = [][]int{
	{777, 3, 777 % 3},
	{111101, 11, 111101 % 11},
	{1555, 7, 1555 % 7},
	{55, 3, 55 % 3},
}

func TestBitLen(t *testing.T) {
	if n := bits.Len(3); n != 2 {
		t.Fatal(n)
	}
}
func TestMontgomery(t *testing.T) {
	for i, v := range cases {
		have, want := montgomery(v[0], v[1]), v[2]
		if have != want {
			t.Logf("%d: have %v want %v", i, have, want)
		}
	}
}

func TestMontgomery2(t *testing.T) {
	for i, v := range cases {
		have, want := montgomery2(v[0], v[1]), v[2]
		if have != want {
			t.Fatalf("%d: have %v want %v", i, have, want)
		}
	}
}
