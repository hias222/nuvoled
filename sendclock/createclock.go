package sendclock

import (
	"fmt"
	"image"
	"image/color"
)

var departure_time = 55
var target_time = [5]int{32, 33, 33, 40, 40}

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

	var minutes int = int(second / 60)
	var seconds_cleared int = second - minutes*60

	var round int = second/departure_time + 1
	//var startsecond_real int = -(minutes*60 - ((round - 1) * departure_time))
	// erst wenn alle weg sind nächste Runde
	var round_start int = int((second - len(target_time)*5) / departure_time)
	var startsecond int = round_start*departure_time - minutes*60
	var swimmerstart int = startsecond + departure_time

	if seconds_cleared > swimmerstart {
		var add int = int((second-swimmerstart)/5+1) * 5
		correct := fmt.Sprintf("correct %d", add)
		logger.Debug(correct)
		swimmerstart = swimmerstart + add
	}

	print := fmt.Sprintf("%d: second %d (%d) start %d swimmerstart %d ", round, second, seconds_cleared, startsecond, swimmerstart)
	logger.Debug(print)

	createImageRelativRGBA(myImg, seconds_cleared, startsecond, round, swimmerstart)
}

func createImageRelativRGBA(myImg *image.RGBA, second int, startsecond int, round int, swimmerstart int) {
	//var colBLUE = color.RGBA{0, 0, 255, 255}
	var colWhite = color.RGBA{255, 255, 255, 255}
	var colGreen = color.RGBA{0, 50, 0, 255}
	var colEnd = color.RGBA{255, 0, 0, 255}

	drawTargetArea(myImg, colGreen, second, startsecond)

	// start next
	addsecondspointer(myImg, swimmerstart, colEnd)

	// zeiger
	addsecondspointer(myImg, second, colWhite)

}
