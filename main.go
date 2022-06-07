package main

import (
	"fmt"

	mqtttoudpclient "swimdata.de/nuvoled/mqttToUdpClient"
	"swimdata.de/nuvoled/mqttclient"
	"swimdata.de/nuvoled/udpserver"
)

func main() {
	//Start UDP
	fmt.Println(udpserver.TestMe())
	s := udpserver.GetLocalAddress()
	connection := udpserver.GetServerClient(s)

	// start MQTT
	var c = mqttclient.IntClientMqtt()
	mqtttoudpclient.IntUDPAddress(s)
	mqttclient.StartCLientMqtt(c)

	udpserver.StartServer(connection)
}
