package main

import (
	"flag"
	"fmt"

	"swimdata.de/nuvoled/mqttclient"
	"swimdata.de/nuvoled/udpserver"
)

func main() {

	bcPtr := flag.String("bc", "192.168.178.255", "broadcast address")
	flag.Parse()
	fmt.Println("-bc", *bcPtr)

	//Start
	fmt.Println(udpserver.TestMe())

	// start MQTT
	var c = mqttclient.IntClientMqtt()
	mqttclient.StartCLientMqtt(c)

	//UDP
	udpserver.InitLocalUdpAdress(*bcPtr)
	udpserver.StartServer()
}
