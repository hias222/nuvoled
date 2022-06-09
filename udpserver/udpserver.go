package udpserver

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"

	"swimdata.de/nuvoled/udpmessages"
)

var udpSource *net.UDPAddr
var udpDestination *net.UDPAddr

func TestMe() string {
	return "Start Test Call .."
}

func InitLocalUdpAdress() {
	PORT := ":2000"

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
	}
	udpSource = s
	fmt.Println("Local Listener Address: ", s.String())

	SENDERPORT := "10.10.10.255:2000"
	sender, err := net.ResolveUDPAddr("udp4", SENDERPORT)
	if err != nil {
		fmt.Println(err)
	}
	udpDestination = sender
	fmt.Println("UDP Detination Address: ", sender.String())

}

func SendUDPMessage(data []byte) {
	c, err := net.DialUDP("udp4", nil, udpDestination)
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()
	_, err = c.Write(data)
}

func handleBufferData(buffer []byte, n int, addr net.Addr) {
	udpmessages.BufferToString(buffer, n)

	if strings.TrimSpace((string(buffer[0 : n-1]))) == "STOP" {
		fmt.Println("Exiting UDP server")
		return
	}

	if n > 3 && buffer[2] == 15 {
		fmt.Print("-> ", string(addr.String()), "\n")
		fmt.Println("Send Messages to panel ")
		SendUDPMessage(udpmessages.CreateRegisterMessage(buffer))
		time.Sleep(1 * time.Second)
		SendUDPMessage(udpmessages.ActivatePanles(buffer))
		time.Sleep(1 * time.Second)
		SendUDPMessage(udpmessages.TurnOnPanles(buffer))

		fmt.Println("Finish Registration")
	}
}

func StartServer() {

	connection, err := net.ListenUDP("udp4", udpSource)
	net.ListenUDP("udp4", udpSource)

	connection.SetReadBuffer(1048576)

	if err != nil {
		fmt.Println(err)
	}

	//defer connection.Close()

	buffer := make([]byte, 2048)

	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println(err)
			return
		}

		go handleBufferData(buffer, n, addr)

	}

}
