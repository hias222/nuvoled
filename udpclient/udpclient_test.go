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

	data := udpmessages.CreateInitMessage()

	testdata := udpclient.SendClientBCInit(data)

	if testdata != "success" {
		t.Fatal("error multicast \n", testdata, "\n")
	}

}
