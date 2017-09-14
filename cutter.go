package main

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"os"

	"github.com/oliamb/cutter"
)

func Cutter() error {
	so, err := os.Open(SS)
	if err != nil {
		return err
	}
	defer so.Close()
	buf := new(bytes.Buffer)
	io.Copy(buf, so)

	img, _, err := image.Decode(buf)
	if err != nil {
		return err
	}
	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:   480,
		Height:  480,
		Anchor:  image.Point{160, 150},
		Options: cutter.Copy,
	})
	out, _ := os.Create(TRIM)
	defer out.Close()
	png.Encode(out, croppedImg)

	return nil
}
