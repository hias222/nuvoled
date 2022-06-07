package udpserver

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"

	"swimdata.de/nuvoled/udpmessages"
)

func TestMe() string {
	return "Start Test Call .."
}

func GetLocalAddress() *net.UDPAddr {
	PORT := "10.10.10.255:2000"

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return s
}

func GetServerClient(s *net.UDPAddr) *net.UDPConn {

	connection, err := net.ListenUDP("udp4", s)
	net.ListenUDP("udp4", s)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(s.String())

	//defer connection.Close()

	return connection
}

func StartServer(connection *net.UDPConn) {

	buffer := make([]byte, 2048)

	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println(err)
			return
		}

		udpmessages.BufferToString(buffer, n)

		if strings.TrimSpace((string(buffer[0 : n-1]))) == "STOP" {
			fmt.Println("Exiting UDP server")
			return
		}

		if n > 3 && buffer[2] == 15 {
			fmt.Print("-> ", string(addr.String()), "\n")
			fmt.Println("Send Messages to panel ")
			_, err = connection.WriteToUDP(udpmessages.CreateRegisterMessage(buffer), addr)
			time.Sleep(1 * time.Second)
			_, err = connection.WriteToUDP(udpmessages.ActivatePanles(buffer), addr)
			time.Sleep(1 * time.Second)
			_, err = connection.WriteToUDP(udpmessages.TurnOnPanles(buffer), addr)
			time.Sleep(1 * time.Second)
			fmt.Println("Finish Registartion")
		}

		data := []byte("hello")
		// fmt.Printf("data: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr)

		if err != nil {
			fmt.Println(err)
			fmt.Println(addr)
			//return
		}

	}

}
