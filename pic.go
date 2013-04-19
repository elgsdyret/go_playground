package main

import "code.google.com/p/go-tour/pic"


func Conversion(x, y int) uint8 {
	return uint8((x*y)/2)	
}

func Pic(dx, dy int) [][]uint8 {
	// each element has length dx
	// there is a number of entries dy 
	yAxis := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		// create the xAxis element
		xAxis := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			xAxis[j] = Conversion(j, i)	
		}
		yAxis[i] = xAxis
	}
	return yAxis
}

func main() {
    pic.Show(Pic)
}
