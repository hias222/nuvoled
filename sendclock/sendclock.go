package sendclock

import (
	"image"
	"image/draw"

	"swimdata.de/nuvoled/logging"
	mqtttoudpclient "swimdata.de/nuvoled/mqttToUdpClient"
)

var logger = logging.GetLogger()

var framenumber int
var baseImage image.RGBA

func SendClock(second int) {

	framenumber++
	// add frame
	if framenumber > 255 {
		framenumber = 1
	}

	logger.Debug("Clock start " + string(second))

	if len(baseImage.Pix) == 0 {
		baseImage = createBaseImage()
		createBackgroundClock(&baseImage)
	}

	var byteRGBA_clock_1 = createBaseImage()
	copy(byteRGBA_clock_1.Pix, baseImage.Pix)
	createImageRGBA(&byteRGBA_clock_1, second)

	var byteRGBA_clock_2 = createBaseImage()
	copy(byteRGBA_clock_2.Pix, baseImage.Pix)
	createImageRGBA(&byteRGBA_clock_2, second)

	rgba := image.NewRGBA(image.Rect(0, 0, 128, 256))

	draw.Draw(rgba, byteRGBA_clock_1.Bounds(), &byteRGBA_clock_1, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, image.Rect(0, 128, 128, 256), &byteRGBA_clock_2, image.Point{0, 0}, draw.Src)

	mqtttoudpclient.SendUDPData(rgba.Pix, framenumber)

}
