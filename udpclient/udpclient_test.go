package udpclient_test

import (
	"testing"

	"swimdata.de/nuvoled/udpclient"
	"swimdata.de/nuvoled/udpmessages"
)

func TestUdpUnicast(t *testing.T) {

	data := udpmessages.CreateInitMessage()

	if udpclient.SendClientInit(data) != "success" {
		t.Fatal("error unicast\n")
	}

}

func TestUdpMulticast(t *testing.T) {

	SERVER := "192.168.178.255"
	//SERVER := "169.254.255.255"

	data := udpmessages.CreateInitMessage()

	testdata := udpclient.SendClientBCInit(data, SERVER)

	if testdata != "success" {
		t.Fatal("error multicast \n", testdata, "\n")
	}

}

/*
Antwort Master->Panel nac
36 $
36 $
160
23
49 1
74 J


Antwort Klick auf Panel
36 $
36 $
120 x
2
32
8
8
1
23
49 1
74 J
0
29
0
0


*/
