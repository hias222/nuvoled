package sendclock

import (
	"image/color"
	"image/draw"
	"math"
)

func konvertPolarToKartDegree(p *pointervalues, degree int, length int) {
	// x = r cos θ , y = r sin θ
	// damit wir oben anfangen
	degreecorrect := degree - 90

	angle := (float64(degreecorrect) / 180 * math.Pi)

	x1 := int(float64(length) * math.Cos(angle))
	y1 := int(float64(length) * math.Sin(angle))

	p.endx = x1 + p.startx
	p.endy = y1 + p.starty
}

func addDegreePointer(img draw.Image, degree int, col color.Color) {

	//a := float64(dy) / float64(dx)
	startx := img.Bounds().Max.X / 2
	starty := img.Bounds().Max.Y / 2

	length := startx - 4

	newpointer := pointervalues{
		startx: startx,
		starty: starty,
		endx:   0,
		endy:   0,
	}

	konvertPolarToKartDegree(&newpointer, degree, length)

	DrawLine(img, newpointer.startx, newpointer.starty, newpointer.endx, newpointer.endy, col)

}

func drawTargetArea(img draw.Image, c_sec color.Color, second int, startsecond int) {

	var count_i = 0
	var cleared_seconds = second - startsecond
	if cleared_seconds < 0 {
		cleared_seconds += 60
	}

	for i := 0; i < len(target_time); i++ {
		var switch_time = target_time[i] + 5*count_i + 2
		if cleared_seconds > switch_time {
			count_i = i
			logger.Debug("step forward " + string(second) + " number " + string(count_i))
		}
	}

	var target_draw = target_time[count_i] + 5*count_i + startsecond

	var calcdegree_start = (target_draw - 1) * 6
	var calcdegree_end = (target_draw + 1) * 6

	for i := calcdegree_start - 1; i < calcdegree_end; i++ {
		addDegreePointer(img, i, c_sec)
	}

}
