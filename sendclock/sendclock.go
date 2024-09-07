package sendclock

import (
	"fmt"

	"swimdata.de/nuvoled/image"
	mqtttoudpclient "swimdata.de/nuvoled/mqttToUdpClient"
)

var framenumber int

var details bool

func SetParameter(details bool) {
	details = details
}

func SendClock() {

	framenumber++
	// add frame
	if framenumber > 255 {
		framenumber = 1
	}

	fmt.Println("Clock start")
	var byteRGBA = image.GetInitImageRGBA()
	mqtttoudpclient.SendUDPData(byteRGBA, framenumber)

}
