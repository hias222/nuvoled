package main

import (
	"fmt"

	"swimdata.de/nuvoled/udpserver"
)

func main() {
	fmt.Println(udpserver.TestMe())

	connection := udpserver.GetServerClient()

	udpserver.StartServer(connection)
}
