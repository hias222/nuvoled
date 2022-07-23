package main

import (
	"flag"
	"fmt"
	"time"

	"swimdata.de/nuvoled/mqttclient"
	"swimdata.de/nuvoled/udpserver"
)

func main() {

	bcPtr := flag.String("bc", "169.254.255.255", "broadcast address")
	ipPtr := flag.Bool("ip", true, "local ip address")
	regPtr := flag.Bool("reg", false, "broadcast address")
	mqttSrv := flag.String("mqtt", "localhost", "mqtt server name")

	flag.Parse()

	//Start
	fmt.Println(udpserver.StartMessage(*bcPtr, *regPtr, *ipPtr, *mqttSrv))

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

	for i := 0; i < 30; i++ {
		fmt.Println("Init Mqtt ... ")
		var c, err = mqttclient.IntClientMqtt(*mqttSrv)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Waiting 30s")
			time.Sleep(30 * time.Second)
		} else {
			fmt.Println("Waiting on mqtt ... ")
			mqttclient.StartCLientMqtt(c)
			fmt.Println("MQTT connected ")
			break
		}
	}

	//UDP
	//go mqtttoudpclient.SendUDPStartMessage()

	fmt.Println("Starting UDP")
	udpserver.InitLocalUdpAdress(*bcPtr, *regPtr, locallistenAddr)
	udpserver.StartServer()

}
