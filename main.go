package main

import (
	"flag"
	"fmt"

	"swimdata.de/nuvoled/mqttclient"
	"swimdata.de/nuvoled/udpserver"
)

func main() {

	ipPtr := flag.String("ip", "192.168.178.175", "local ip")
	bcPtr := flag.String("bc", "192.168.178.255", "broadcast address")
	flag.Parse()
	fmt.Println("-ip", *ipPtr)
	fmt.Println("-bc", *bcPtr)

	if *ipPtr == "" {
		fmt.Println("missing")
	}

	//Start
	fmt.Println(udpserver.TestMe())

	// start MQTT
	var c = mqttclient.IntClientMqtt()
	mqttclient.StartCLientMqtt(c)

	//UDP
	udpserver.InitLocalUdpAdress(*ipPtr, *bcPtr)
	udpserver.StartServer()
}
