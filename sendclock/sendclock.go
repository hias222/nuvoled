package sendclock

import (
	"fmt"

	mqtttoudpclient "swimdata.de/nuvoled/mqttToUdpClient"
)

var framenumber int

func SendClock() {

	framenumber++
	// add frame
	if framenumber > 255 {
		framenumber = 1
	}

	fmt.Println("Clock start")

	var byteRGBA = createImageRGBA()
	mqtttoudpclient.SendUDPData(byteRGBA, framenumber)

}
