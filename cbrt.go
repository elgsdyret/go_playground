package main

import (
	"fmt"
	"math/cmplx"
)

//z = z - (z^3-x)/(3*z^2)
func Newton(x, z complex128) complex128 {
	return z - ((cmplx.Pow(z, 3) - x) / (3* (cmplx.Pow(z, 2))))
}

func Cbrt(x complex128) complex128 {
		
	z := complex128(1)
	for i := 0; i < 10000; i++ {
		z = Newton(z, x)
	}
	return z
}

func main() {
    fmt.Println(Cbrt(2))
    fmt.Println(cmplx.Pow(2, (1/3)))
}
