package udpserver_test

/*
36 $
36 $
15
0
74 J
23
49 1
80 P
52 4
83 S
32
8
8
8
8
-> $$J1P4S
MFU@QM-MOS-MFU nuvoled % sudo arp 169.254.23.49
? (169.254.23.49) at 12:80:e1:4a:17:31 on en6 [ethernet]
*/

import (
	"fmt"
	"testing"

	"swimdata.de/nuvoled/udpserver"
)

func TestUdpserver(t *testing.T) {

	bcPtr := "169.254.255.255"
	ipPtr := true       // "local ip address")
	regPtr := false     //, "broadcast address")
	mqttSrv := "ubuntu" //", "mqtt server name")

	fmt.Println(udpserver.StartMessage(bcPtr, regPtr, ipPtr, mqttSrv))

	locallistenAddr := "127.0.0.1"

	udpserver.InitLocalUdpAdress(bcPtr, regPtr, locallistenAddr)
	fmt.Println("Starting Server")
	udpserver.StartServer()
	//udpserver.StartSimple()

	t.Fatal("test")

}
