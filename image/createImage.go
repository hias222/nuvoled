package image

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"swimdata.de/nuvoled/logging"
)

var logger = logging.GetLogger()

var (
	dpiImage      = flag.Float64("dpiImage", 144, "screen resolution in Dots Per Inch")
	fontfileImage = flag.String("fontfileImage", "static/fonts/fixed_bold.ttf", "filename of the ttf font")
	sizeImage     = flag.Float64("sizeImage", 16, "font size in points")
	debug         = flag.Bool("debug", false, "debug mode")
)

//go:embed static/fonts/fixed_bold.ttf
//go:embed base/logo_128_128.png
var embedContent embed.FS

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

func addLabelFont(img *image.RGBA, x, y int, top string, button string) {

	flag.Parse()

	logger.Debug("addLabelFont")

	// Read the font data from embed
	fontBytes, err := embedContent.ReadFile(*fontfileImage)
	//fontBytes, err := ioutil.ReadFile(mebedContent)
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

	fgColor := color.RGBA{255, 255, 255, 255}
	fg := image.NewUniform(fgColor)
	bgColor := color.RGBA{0, 0, 205, 0xff}
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

	if _, err := c.DrawString(top, pt); err != nil {
		// handle error
		log.Println("Error Draw")
		fmt.Println(err)
		return
	}

	pt = freetype.Pt(x, y+64+int(c.PointToFixed(*sizeImage)>>6))

	if _, err := c.DrawString(button, pt); err != nil {
		// handle error
		log.Println("Error Draw")
		fmt.Println(err)
		return
	}

	logger.Debug("End Fonts Draw")
}

func CreateImageRGBA(event string, heat string) []byte {

	myImg := image.NewRGBA(image.Rect(0, 0, 128, 128))

	addLabelFont(myImg, 3, 15, event, heat)

	//udpmessages.BufferToString(myImg.Pix, 10024)

	if *debug {
		fmt.Printf("myImg: %v\n", myImg)
		out, err := os.Create("cat.png")
		png.Encode(out, myImg)
		out.Close()

		if err != nil {
			fmt.Println(err)
		}
	}
	return myImg.Pix
}

func GetInitImageRGBA() []byte {

	infile, err := embedContent.ReadFile("base/logo_128_128.png")

	if err != nil {
		// replace this with real error handling
		return nil
	}

	src, err := png.Decode(bytes.NewReader(infile))
	if err != nil {
		// replace this with real error handling
		return nil
	}

	b := src.Bounds()
	myImg := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(myImg, myImg.Bounds(), src, b.Min, draw.Src)

	if *debug {
		fmt.Printf("myImg: %v\n", myImg)
		out, err := os.Create("start.png")
		png.Encode(out, myImg)
		out.Close()

		if err != nil {
			fmt.Println(err)
		}
	}

	return myImg.Pix
}
