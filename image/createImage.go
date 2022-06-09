package image

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

var (
	dpiImage      = flag.Float64("dpiImage", 72, "screen resolution in Dots Per Inch")
	fontfileImage = flag.String("fontfileImage", "../static/fonts/UbuntuMono-R.ttf", "filename of the ttf font")
	sizeImage     = flag.Float64("sizeImage", 32, "font size in points")
)

//https://stackoverflow.com/questions/38299930/how-to-add-a-simple-text-label-to-an-image-in-go

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func addLabelFont(img *image.RGBA, x, y int, label string) {

	flag.Parse()

	log.Println("addLabelFont")

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(*fontfileImage)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	//fg := image.Black

	fgColor := color.RGBA{0, 0, 0, 255}
	fg := image.NewUniform(fgColor)
	bgColor := color.RGBA{2, 255, 2, 0xff}
	bg := image.NewUniform(bgColor)
	//bg := image.Black

	//draw.Draw(img, img.Bounds(), bg, image.ZP, draw.Src)
	draw.Draw(img, img.Bounds(), bg, image.Pt(0, 0), draw.Src)

	c := freetype.NewContext()
	c.SetDPI(*dpiImage)
	c.SetFont(f)
	c.SetFontSize(*sizeImage)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(fg)
	c.SetDst(img)
	//size := 12.0 // font size in pixels
	pt := freetype.Pt(x, y+int(c.PointToFixed(*sizeImage)>>6))

	if _, err := c.DrawString(label, pt); err != nil {
		// handle error
		log.Println("Error Draw")
		fmt.Println(err)
		return
	}
	log.Println("End Fonts Draw")
}

func CreateImage() {

	myImg := image.NewRGBA(image.Rect(0, 0, 128, 128))

	addLabelFont(myImg, 0, 0, "Hello")

	//udpmessages.BufferToString(myImg.Pix, 10024)

	fmt.Printf("myImg: %v\n", myImg)

	out, err := os.Create("cat.png")
	png.Encode(out, myImg)
	out.Close()

	if err != nil {
		fmt.Println(err)
	}
}
