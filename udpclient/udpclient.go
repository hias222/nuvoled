package udpclient

import (
	"fmt"
	"net"
)

func SendClientInit(data []byte) string {

	CONNECT := "127.0.0.1:2000"
	//MESSAGE := "$$J1P4S"

	fmt.Println("conect to ", CONNECT)

	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return "error"
	}

	fmt.Println("the UDP Server is ", c.RemoteAddr().String())

	defer c.Close()

	_, err = c.Write(data)

	buffer := make([]byte, 1024)
	n, _, err := c.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		return "error"
	}

	fmt.Println("Reply:", string(buffer[0:n]))

	if string(buffer[0:n]) != "hello" {
		fmt.Println("error")
		return "error"
	}
	return "success"
}