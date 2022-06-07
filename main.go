package main

import (
	"fmt"

	"swimdata.de/nuvoled/mqttclient"
	"swimdata.de/nuvoled/udpserver"
)

func main() {
	//Start
	fmt.Println(udpserver.TestMe())

	// start MQTT
	var c = mqttclient.IntClientMqtt()
	mqttclient.StartCLientMqtt(c)

	//UDP
	udpserver.InitLocalUdpAdress()
	udpserver.StartServer()
}
