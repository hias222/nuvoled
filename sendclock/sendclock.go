package sendclock

import (
	"fmt"

	mqtttoudpclient "swimdata.de/nuvoled/mqttToUdpClient"
)

var framenumber int

func SendClock(second int) {

	framenumber++
	// add frame
	if framenumber > 255 {
		framenumber = 1
	}

	fmt.Printf("Clock start % d ", second)

	var byteRGBA = createImageRGBA(second)
	mqtttoudpclient.SendUDPData(byteRGBA, framenumber)

}
