
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Newton(z, x float64) float64 {
	return z - ((math.Pow(z, 2) - x) / (2*z))	
}

// should return non nil error when negative - as it does not support complex numbers
func Sqrt(x float64) (float64, error) {
	
	if x <= 0 {
		return math.NaN(), ErrNegativeSqrt(x)
	}

	z, zNext := 1.0, 1.0
	
	for i := 0; i < 100; i++ {
		z = Newton(z, x)
		zNext = Newton(z, x)
		
		delta := z - zNext 
		if delta < 0.00000001 {
			fmt.Println("iteration", i)
			return z, nil
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(-2))
	fmt.Println(math.Sqrt(-2))
}

