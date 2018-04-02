package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"time"

	"github.com/golang/tour/pic"
)

type Image struct {
	W int
	H int
}

func (self Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, self.W, self.H)
}

func (self Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (self Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	var buf bytes.Buffer
	m := Image{W: 100, H: 100}
	pic.ShowImage(m)
	png.Encode(&buf, m)
	ioutil.WriteFile("a.png", buf.Bytes(), os.FileMode.Perm(os.ModeAppend))
	time.Sleep(10 * time.Second)
}
