package udpclient

import (
	"fmt"
	"net"
)

func SendClientBCInit(data []byte, server string) string {

	PORT := ":2000"
	CONNECT := server + PORT

	fmt.Println("conect to ", CONNECT)

	pc, err := net.ListenPacket("udp4", PORT)

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

	fmt.Println("Resolve Server is ", s.String())

	_, err = pc.WriteTo(data, s)

	if err != nil {
		fmt.Println(err)
		//fmt.Println(i)
		return "error WriteTo"
	}

	/*
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
	*/
	return "success"
}
