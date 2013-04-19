package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

func Conversion(x, y int) uint8 {
	return uint8((x*y)/2)	
}

type Image struct{
	Width	int
    Height	int
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.Width, i.Height)    
}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
    v := Conversion(x, y)
    return color.RGBA{v, v, 255, 255}
}

func main() {
    m := Image { 200, 200}
    pic.ShowImage(m)
}
