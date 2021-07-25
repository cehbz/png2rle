package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func bw(c color.Color) bool {
	return color.NRGBAModel.Convert(c).(color.NRGBA).A > 128
}

func main() {
	f, err := os.Open(os.Args[1])
	checkErr(err)
	img, err := png.Decode(f)
	checkErr(err)
	b := img.Bounds()
	curColor := false
	runLen := 0
	fmt.Printf("const uint8_t rle[] = {%d,%d\n", b.Dx(), b.Dy())
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := bw(img.At(x, y))
			if curColor == c {
				runLen++
				if runLen > 255 {
					fmt.Print(",255,0")
					runLen -= 255
				}
			} else {
				fmt.Printf(",%d", runLen)
				runLen = 0
				curColor = c
			}
		}
	}
	if runLen > 0 {
		fmt.Printf(",%d", runLen)
	}
	fmt.Println("};")
}
