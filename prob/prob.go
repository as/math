package main

import (
	"fmt"
)

type A []int

func (a A) And(b A) A {
	c := A{}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			c = append(c, a[i]&b[j])
		}
	}
	return c
}
func (a A) Or(b A) A {
	c := A{}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			c = append(c, a[i]|b[j])
		}
	}
	return c
}
func (a A) Xor(b A) A {
	c := A{}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			c = append(c, a[i]^b[j])
		}
	}
	return c
}
func (a A) Not() A {
	c := A{}
	for i := 0; i < len(a); i++ {
		c = append(c, (a[i]+1)&1)
	}
	return c
}
func (a A) String() string {
	n := 0
	for _, x := range a {
		if x == 1 {
			n++
		}
	}
	m := len(a)
	if m != 16 {

		g := gcd(n, m)
		if g != 1 {
			n /= g
			m /= g
		}

		if m == 2 {
			m *= 2
			n *= 2
		}
		if m == 4 {
			m *= 2
			n *= 2
		}
		if m == 8 {
			m *= 2
			n *= 2
		}
	}
	return fmt.Sprintf("%02d/%02d %v", n, m, []int(a))
}
func half(n int) A {
	a := make(A, n)
	for i := 0; i < n; i++ {
		a[i] = i & 1
	}
	return a

}
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
func main() {
	x := half(2)
	z := x.And(x)
	p08 := half(16)
	p07 := z.Or(z)
	p06 := z.Xor(z)
	p12 := half(4).Or(half(4)) // p04.Not()
	p04 := half(4).And(half(4))
	p10 := p04.Xor(p12)
	p05 := p08.And(p10)
	p03 := p04.And(p12)
	p01 := p04.And(p04)
	p13 := p04.Or(p12)
	p14 := p08.Or(p12)
	p11 := p05.Not()
	p09 := p07.Not()
	p02 := p14.Not()
	p15 := p01.Not()

	fmt.Println(p01)
	fmt.Println(p02)
	fmt.Println(p03)
	fmt.Println(p04)
	fmt.Println(p05)
	fmt.Println(p06)
	fmt.Println(p07)
	fmt.Println(p08)
	fmt.Println(p09)
	fmt.Println(p10)
	fmt.Println(p11)
	fmt.Println(p12)
	fmt.Println(p13)
	fmt.Println(p14)
	fmt.Println(p15)
}
