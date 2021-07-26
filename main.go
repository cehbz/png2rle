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
	s := os.Args[1][:len(os.Args[1])-4]
	f, err := os.Open(s + ".png")
	checkErr(err)
	img, err := png.Decode(f)
	checkErr(err)
	b := img.Bounds()
	curColor := false
	runLen := 0
	sep := ""
	fmt.Printf("const uint8_t bitmap_%s[] = {\n", s)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := bw(img.At(x, y))
			if curColor == c {
				runLen++
				if runLen > 255 {
					fmt.Print(sep, "255,0")
					sep = ","
					runLen -= 255
				}
			} else {
				fmt.Printf("%s%d", sep, runLen)
				sep = ","
				runLen = 1
				curColor = c
			}
		}
	}
	if runLen > 0 {
		fmt.Printf("%s%d", sep, runLen)
		sep = ","
	}
	fmt.Println("};")
	fmt.Printf("const rle rle_%s = {%d, %d, bitmap_%s};\n", s, b.Dx(), b.Dy(), s)
	fmt.Println()
}
