package udpserver_test

import (
	"net"
	"testing"
)

func TestUdpserver(t *testing.T) {
	//if udpserver.TestMe() != "test" {
	//	t.Fatal("error")
	//}

	CONNECT := "127.0.0.1:2000"

	t.Log("conect to ", CONNECT)

	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("the UDP Server is ", c.RemoteAddr().String())

	defer c.Close()

	data := []byte("STOP\n")

	_, err = c.Write(data)

	if err != nil {
		t.Fatal(err)
	}

	t.Fatal("Ende")

}
