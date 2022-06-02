package udpclient

import (
	"fmt"
	"net"
)

func SendClientBCInit(data []byte) string {

	CONNECT := "169.254.154.141:2000"
	fmt.Println("conect to ", CONNECT)

	pc, err := net.ListenPacket("udp4", ":2000")
	if err != nil {
		fmt.Println(err)
		return "error ListenPacket"
	}

	defer pc.Close()

	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	if err != nil {
		fmt.Println(err)
		return "error ResolveUDPAddr"
	}

	i, err := pc.WriteTo(data, s)

	if err != nil {
		fmt.Println(err)
		fmt.Println(i)
		return "error WriteTo"
	}

	buffer := make([]byte, 1024)
	n, _, err := pc.ReadFrom(buffer)
	if err != nil {
		fmt.Println(err)
		return "error ReadFrom"
	}

	answer := string(buffer[0:n])
	fmt.Println("Reply:", answer)

	if answer != "hello" {
		return "error got " + answer
	}

	return "success"
}
