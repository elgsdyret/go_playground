package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := 0
	fibNMinus1 := 1
	fibNMinus2 := 0

	return func() int {
		if i == 0  {
			i++
			return 0
		}
		if (i == 1) {
			i++
			return 1
		}
		newFib := fibNMinus1 + fibNMinus2
		fibNMinus2 = fibNMinus1
		fibNMinus1 = newFib
		return newFib
	}
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
