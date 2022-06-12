package main

import (
	"flag"
	"fmt"

	"swimdata.de/nuvoled/mqttclient"
	"swimdata.de/nuvoled/udpserver"
)

func main() {

	bcPtr := flag.String("bc", "169.254.255.255", "broadcast address")
	regPtr := flag.Bool("reg", true, "broadcast address")
	flag.Parse()

	//Start
	fmt.Println(udpserver.StartMessage(*bcPtr, *regPtr))

	// start MQTT
	var c = mqttclient.IntClientMqtt()
	mqttclient.StartCLientMqtt(c)

	//UDP
	udpserver.InitLocalUdpAdress(*bcPtr, *regPtr)
	udpserver.StartServer()
}
