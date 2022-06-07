package mqtttoudpclient

import (
	"fmt"
	"net"
)

var udpSource *net.UDPAddr

func IntUDPAddress(laddr *net.UDPAddr) {
	udpSource = laddr
}

func SendUDPMessage(data []byte) {
	c, err := net.DialUDP("udp4", nil, udpSource)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("the UDP Server is ", c.RemoteAddr().String())

	defer c.Close()

	_, err = c.Write(data)
}
