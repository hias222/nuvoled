package sendclock

import (
	"image"

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

	var byteRGBA2 = createBaseImage()
	copy(byteRGBA2.Pix, baseImage.Pix)

	createImageRGBA(&byteRGBA2, second)

	mqtttoudpclient.SendUDPData(byteRGBA2.Pix, framenumber)

}
