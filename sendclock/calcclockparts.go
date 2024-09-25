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

func addsecondspointer(img draw.Image, second int) {

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

	var colBLUE = color.RGBA{0, 0, 255, 255}
	DrawLine(img, newpointer.startx, newpointer.starty, newpointer.endx, newpointer.endy, colBLUE)

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
