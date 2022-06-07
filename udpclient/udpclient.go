package udpclient

import (
	"fmt"
	"net"
)

func SendClientInit(data []byte) string {

	//CONNECT := "127.0.0.1:2000"
	//169.254.255.255
	//CONNECT := "192.168.178.255:2000"

	//CONNECT := "169.254.255.255:2000"
	//fmt.Println("conect to ", CONNECT)

	s, err := net.ResolveUDPAddr("udp4", "10.10.10.7:2000")

	fmt.Println("Resolve Server is ", s.String())

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
