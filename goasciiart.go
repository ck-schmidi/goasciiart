// by mo2zie
// changes by ck-schmidi

package goasciiart

import (
	"github.com/nfnt/resize"

	"bytes"
	"image"
	"image/color"
	"reflect"
)

// ASCIISTR are the characters used for creating ascii image
var ASCIISTR = "MND8OZ$7I?+=~:,.."

// ScaleImage scales a given image the a given width and
// returns the image + the new size (width, height)
func ScaleImage(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
	return img, w, h
}

// Convert2Ascii converts a given image with size to an ascii image
func Convert2Ascii(img image.Image, w, h int) []byte {
	table := []byte(ASCIISTR)
	buf := new(bytes.Buffer)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buf.WriteByte(table[pos])
		}
		_ = buf.WriteByte('\n')
	}
	return buf.Bytes()
}
