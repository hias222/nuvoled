package sendclock

import (
	"image"
	"image/color"
)

var departure_time = 30
var target_time = [5]int{13, 14, 15, 16, 16}

//var logger = logging.GetLogger()

func initBaseConfig(departure int, target [5]int) {
	departure_time = departure
	target_time = target
}

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

	bgColor := color.RGBA{0, 0, 255, 0xff}
	bgColor5 := color.RGBA{0, 255, 255, 0xff}

	drawCircle(myImg, bgColor)
	drawLiniesOfCircle(myImg, bgColor, bgColor5)

}

func createImageRGBA(myImg *image.RGBA, second int) {
	//var colBLUE = color.RGBA{0, 0, 255, 255}
	var colWhite = color.RGBA{255, 255, 255, 255}
	var colGreen = color.RGBA{0, 50, 0, 255}
	var colEnd = color.RGBA{255, 0, 0, 255}

	startsecond := 0

	drawTargetArea(myImg, colGreen, second, startsecond)

	// start next
	var start_time = departure_time + startsecond
	if second > departure_time+startsecond {
		var diff = second - departure_time - startsecond + 5
		var anzahl = diff / 5
		start_time += 5 * anzahl
	}
	addsecondspointer(myImg, start_time, colEnd)

	// zeiger
	addsecondspointer(myImg, second, colWhite)

}
