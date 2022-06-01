package udpserver

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func TestMe() string {
	return "test"
}

func StartServer() {
	PORT := ":2000"

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()

	buffer := make([]byte, 1024)

	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]), "\n")

		if strings.TrimSpace((string(buffer[0 : n-1]))) == "STOP" {
			fmt.Println("Exiting UDP server")
			return
		}

		if err != nil {
			fmt.Println(err)
			fmt.Println(addr)
			return
		}

	}

}

func random(i1, i2 int) {
	panic("unimplemented")
}
