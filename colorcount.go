package main

import (
	//"fmt"
	"image/color"
	"image/png"
	"os"
)

var (
	r1   = color.RGBA{153, 255, 255, 0}
	r5   = color.RGBA{103, 205, 255, 0}
	r10  = color.RGBA{30, 146, 255, 0}
	r20  = color.RGBA{0, 56, 255, 0}
	r30  = color.RGBA{250, 245, 0, 0}
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

func rgbToInteger(r, g, b uint8) (i int) {
	i = (int)(r)
	i <<= 8
	i += (int)(g)
	i <<= 8
	i += (int)(b)
	return
}

func limSet(v uint8, i int) uint8 {
	if (int)(v)+i < 0 {
		return 0
	} else if (int)(v)+i > 255 {
		return 255
	} else {
		return v + (uint8)(i)
	}
}

func appColSearch(c color.RGBA, rx color.RGBA) bool {
	r := rx.R
	g := rx.G
	b := rx.B
	ci := toInteger(c)
	//fmt.Printf("R:%d G:%d B:%d\n", c.R, c.G, c.B)

	for i := -APP_RANGE; i <= APP_RANGE; i++ {
		for j := -APP_RANGE; j <= APP_RANGE; j++ {
			for k := -APP_RANGE; k <= APP_RANGE; k++ {
				tr := limSet(r, i)
				tg := limSet(g, j)
				td := limSet(b, k)
				xi := rgbToInteger(tr, tg, td)
				if ci == xi {
					return true
				}
			}
		}
	}
	return false
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
			if appColSearch(rgba, r1) {
				w += 1
			} else if appColSearch(rgba, r5) {
				w += 5
			} else if appColSearch(rgba, r10) {
				w += 10
			} else if appColSearch(rgba, r20) {
				w += 20
			} else if appColSearch(rgba, r30) {
				w += 30
			} else if appColSearch(rgba, r50) {
				w += 50
			} else if appColSearch(rgba, r100) {
				w += 100
			}
		}
	}

	return w, nil
}
