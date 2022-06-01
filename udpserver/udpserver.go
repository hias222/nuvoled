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

		i := 0
		for i < n {
			fmt.Print(buffer[i])
			fmt.Print(" ", string(buffer[i]), "\n")
			i++
		}

		fmt.Print("-> ", string(buffer[0:n]), "\n")
		fmt.Print("-> ", string(addr.String()), "\n")

		if strings.TrimSpace((string(buffer[0 : n-1]))) == "STOP" {
			fmt.Println("Exiting UDP server")
			return
		}

		data := []byte("hello")
		fmt.Printf("data: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr)

		if err != nil {
			fmt.Println(err)
			return
		}

		if err != nil {
			fmt.Println(err)
			fmt.Println(addr)
			return
		}

	}

}
