package traffic

import (
	"flag"
	"fmt"

	"swimdata.de/nuvoled/udpserver"
)

func Read(checkIP bool) {

	flag.Parse()

	var locallistenAddr = ""
	if checkIP {
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

	fmt.Println(locallistenAddr)

	fmt.Println("Starting UDP")
	udpserver.StartServer()

}
