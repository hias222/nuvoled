package sendclock

import (
	"image"
	"image/color"
)

func createBaseImage() image.RGBA {

	myImg := image.NewRGBA(image.Rect(0, 0, 128, 128))

	/*
		bgColor := color.RGBA{40, 100, 205, 0xff}
		bg := image.NewUniform(bgColor)

		//myImg := image.NewRGBA(image.Rect(0, 0, 128, 128))

		draw.Draw(myImg, myImg.Bounds(), bg, image.Pt(0, 0), draw.Src)
	*/

	return *myImg
}

func createBackgroundClock(myImg *image.RGBA) {

	//udpmessages.BufferToString(myImg.Pix, 10024)

	bgColor := color.RGBA{0, 0, 255, 0xff}

	/*
		for x := 20; x < myImg.Rect.Dx()-10; x++ {
			myImg.SetRGBA(x, 50, bgColor)
			myImg.SetRGBA(x, 51, bgColor)
		}
	*/

	drawCircle(myImg, bgColor)

}

func createImageRGBA(myImg *image.RGBA, second int) {
	addsecondspointer(myImg, second)
}
