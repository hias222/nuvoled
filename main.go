package main

import (
	"flag"
	"fmt"

	"swimdata.de/nuvoled/mqttclient"
	"swimdata.de/nuvoled/udpserver"
)

func main() {

	bcPtr := flag.String("bc", "169.254.255.255", "broadcast address")
	ipPtr := flag.Bool("ip", true, "local ip address")
	regPtr := flag.Bool("reg", true, "broadcast address")
	flag.Parse()

	//Start
	fmt.Println(udpserver.StartMessage(*bcPtr, *regPtr, *ipPtr))

	var locallistenAddr = ""
	if *ipPtr {
		selfaddress := udpserver.GetIpaddress()

		if selfaddress == "" {
			fmt.Println("no self assigned")
			selfaddress = udpserver.GetLocaladdress()
			if selfaddress == "" {
				fmt.Println("nothing found (192 vs 169")
			}
		}
		fmt.Println(selfaddress)
		locallistenAddr = selfaddress
	}

	// start MQTT
	var c = mqttclient.IntClientMqtt()
	mqttclient.StartCLientMqtt(c)

	//UDP
	udpserver.InitLocalUdpAdress(*bcPtr, *regPtr, locallistenAddr)
	udpserver.StartServer()
}
