package sendclock

import (
	"image"
	"image/color"
	"image/draw"
)

func createImageRGBA() []byte {

	myImg := image.NewRGBA(image.Rect(0, 0, 128, 128))

	bgColor := color.RGBA{40, 100, 205, 0xff}
	bg := image.NewUniform(bgColor)

	draw.Draw(myImg, myImg.Bounds(), bg, image.Pt(0, 0), draw.Src)

	for x := 10; x < myImg.Rect.Dx()-10; x++ {
		myImg.SetRGBA(x, 10, color.RGBA{255, 255, 255, 255})
		myImg.SetRGBA(x, 11, color.RGBA{255, 255, 255, 255})
	}

	var colBLUE = color.RGBA{0, 0, 255, 255}

	// draw line
	DrawLine(myImg, 14, 71, 100, 100, colBLUE)

	//udpmessages.BufferToString(myImg.Pix, 10024)

	return myImg.Pix
}
