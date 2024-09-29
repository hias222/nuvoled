package sendclock

import (
	"image/color"
	"image/draw"
	"math"
)

type pointervalues struct {
	startx int
	starty int
	endx   int
	endy   int
}

func konvertPolarToKart(p *pointervalues, second int, length int) {
	// x = r cos θ , y = r sin θ
	// damit wir oben anfangen
	secondcorrect := second - 15

	angle := (float64(secondcorrect*6) / 180 * math.Pi)

	x1 := int(float64(length) * math.Cos(angle))
	y1 := int(float64(length) * math.Sin(angle))

	p.endx = x1 + p.startx
	p.endy = y1 + p.starty
}

func addsecondspointer(img draw.Image, second int, colBLUE color.Color) {

	//a := float64(dy) / float64(dx)
	startx := img.Bounds().Max.X / 2
	starty := img.Bounds().Max.Y / 2

	length := startx - 5

	newpointer := pointervalues{
		startx: startx,
		starty: starty,
		endx:   0,
		endy:   0,
	}

	konvertPolarToKart(&newpointer, second, length)

	DrawLine(img, newpointer.startx, newpointer.starty, newpointer.endx, newpointer.endy, colBLUE)

}

func drawLiniesOfCircle(img draw.Image, c_sec color.Color, c_5sec color.Color) {

	startx := img.Bounds().Max.X / 2
	starty := img.Bounds().Max.Y / 2

	for i := 0; i < 60; i++ {
		drawLinieOfCircle(img, c_sec, i, 2, startx, starty)
	}

	for i := 0; i < 60; i += 5 {
		drawLinieOfCircle(img, c_5sec, i, 4, startx, starty)
		drawLinieOfCircle(img, c_5sec, i, 4, startx+1, starty)
		drawLinieOfCircle(img, c_5sec, i, 4, startx, starty+1)
		drawLinieOfCircle(img, c_5sec, i, 4, startx+1, starty+1)
	}

}

func drawLinieOfCircle(img draw.Image, c color.Color, seconds int, length int, startx int, starty int) {

	lenth_out := img.Bounds().Max.X/2 - 4
	length_in := img.Bounds().Max.X/2 - 4 - length

	newpointer_in := pointervalues{
		startx: startx,
		starty: starty,
		endx:   0,
		endy:   0,
	}

	newpointer_out := pointervalues{
		startx: startx,
		starty: starty,
		endx:   0,
		endy:   0,
	}

	konvertPolarToKart(&newpointer_in, seconds, length_in)

	konvertPolarToKart(&newpointer_out, seconds, lenth_out)

	DrawLine(img, newpointer_in.endx, newpointer_in.endy, newpointer_out.endx, newpointer_out.endy, c)

}

func drawCircle(img draw.Image, c color.Color) {

	startx := img.Bounds().Max.X / 2
	starty := img.Bounds().Max.Y / 2
	r := startx - 3

	x0 := startx
	y0 := starty

	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)

	for x > y {
		img.Set(x0+x, y0+y, c)
		img.Set(x0+y, y0+x, c)
		img.Set(x0-y, y0+x, c)
		img.Set(x0-x, y0+y, c)
		img.Set(x0-x, y0-y, c)
		img.Set(x0-y, y0-x, c)
		img.Set(x0+y, y0-x, c)
		img.Set(x0+x, y0-y, c)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}
}
