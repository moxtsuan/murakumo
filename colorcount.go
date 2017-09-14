package main

import (
	"image/color"
	"image/png"
	"os"
)

var (
	r1   = color.RGBA{153, 255, 255, 0}
	r5   = color.RGBA{103, 205, 255, 0}
	r10  = color.RGBA{30, 146, 255, 0}
	r20  = color.RGBA{38, 0, 255, 0}
	r30  = color.RGBA{248, 252, 0, 0}
	r50  = color.RGBA{255, 147, 0, 0}
	r80  = color.RGBA{252, 0, 0, 0}
	r100 = color.RGBA{154, 0, 121, 0}
)

func toInteger(rgba color.RGBA) (i int) {
	i = (int)(rgba.R)
	i <<= 8
	i += (int)(rgba.G)
	i <<= 8
	i += (int)(rgba.B)
	return
}

func toRGBA(c color.Color) (rgba color.RGBA) {
	r, g, b, _ := c.RGBA()
	rgba.R = (uint8)(r >> 8)
	rgba.G = (uint8)(g >> 8)
	rgba.B = (uint8)(b >> 8)
	rgba.A = 0
	//fmt.Printf("R:%d G:%d B:%d\n", rgba.R, rgba.G, rgba.B)
	return
}

func ColorCount() (int, error) {
	var w int
	f, err := os.Open(TRIM)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	img, _ := png.Decode(f)
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.RGBAModel.Convert(img.At(x, y))
			rgba := toRGBA(c)
			i := toInteger(rgba)
			switch i {
			case toInteger(r1):
				w += 1
			case toInteger(r5):
				w += 5
			case toInteger(r10):
				w += 10
			case toInteger(r20):
				w += 20
			case toInteger(r30):
				w += 30
			case toInteger(r50):
				w += 50
			case toInteger(r80):
				w += 80
			case toInteger(r100):
				w += 100
			}
		}
	}

	return w, nil
}
