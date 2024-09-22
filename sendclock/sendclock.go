package sendclock

import (
	"swimdata.de/nuvoled/logging"
	mqtttoudpclient "swimdata.de/nuvoled/mqttToUdpClient"
)

var logger = logging.GetLogger()

var framenumber int

func SendClock(second int) {

	framenumber++
	// add frame
	if framenumber > 255 {
		framenumber = 1
	}

	logger.Debug("Clock start " + string(second))

	var byteRGBA = createImageRGBA(second)
	mqtttoudpclient.SendUDPData(byteRGBA, framenumber)

}
