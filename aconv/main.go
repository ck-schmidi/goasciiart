package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/ck-schmidi/goasciiart"
)

func main() {
	w := flag.Int("w", 80, "Use -w <width>")
	fpath := flag.String("p", "test.jpg", "Use -p <filesource>")
	flag.Parse()

	f, err := os.Open(*fpath)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	f.Close()

	img, width, height := goasciiart.ScaleImage(img, *w)

	p := goasciiart.Convert2Ascii(img, width, height)
	fmt.Print(string(p))
}
