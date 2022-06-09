package mqtttoudpclient_test

import (
	"testing"

	mqtttoudpclient "swimdata.de/nuvoled/mqttToUdpClient"
)

func TestSendMQTT(t *testing.T) {
	b := []byte("header 001 001")
	mqtttoudpclient.SendUDPMessage(b)
	t.Fail()
}
