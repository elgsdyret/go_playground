
package main

import (
	"fmt"
	"math"
)

func Newton(z, x float64) float64 {
	return z - ((math.Pow(z, 2) - x) / (2*z))	
}

func Sqrt(x float64) float64 {
	z, zNext := 1.0, 1.0
	
	for i := 0; i < 100; i++ {
		z = Newton(z, x)
		zNext = Newton(z, x)
		
		delta := z - zNext 
		if delta < 0.00000001 {
			fmt.Println("iteration", i)
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1188197))
	fmt.Println(math.Sqrt(1188197))
}
