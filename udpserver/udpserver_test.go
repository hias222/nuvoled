package udpserver_test

/*
36 $
36 $
15
0
74 J
23
49 1
80 P
52 4
83 S
32
8
8
8
8
-> $$J1P4S
MFU@QM-MOS-MFU nuvoled % sudo arp 169.254.23.49
? (169.254.23.49) at 12:80:e1:4a:17:31 on en6 [ethernet]
*/

import (
	"net"
	"testing"
)

func TestUdpserver(t *testing.T) {
	//if udpserver.TestMe() != "test" {
	//	t.Fatal("error")
	//}

	CONNECT := "127.0.0.1:2000"
	MESSAGE := "$$J1P4S"

	t.Log("conect to ", CONNECT)

	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("the UDP Server is ", c.RemoteAddr().String())

	defer c.Close()

	data := []byte(MESSAGE)

	_, err = c.Write(data)

	buffer := make([]byte, 1024)
	n, _, err := c.ReadFromUDP(buffer)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log("Reply:", string(buffer[0:n]))

	if string(buffer[0:n]) != "hello" {
		t.Log("error")
		return
	}
	t.Fatal("test")

}
