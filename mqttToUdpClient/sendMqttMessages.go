package mqtttoudpclient

import (
	"swimdata.de/nuvoled/udpserver"
)

func SendUDPMessage(data []byte) {
	udpserver.SendUDPMessage(data)
}
