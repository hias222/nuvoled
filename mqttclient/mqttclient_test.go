package mqttclient_test

import (
	"testing"
	"time"

	"swimdata.de/nuvoled/mqttclient"
)

func TestMqttClient(t *testing.T) {
	var c = mqttclient.IntClientMqtt("localhost")
	mqttclient.TestClientMqtt(c)
	t.Log("End MQTT")
}

func TestClient15s(t *testing.T) {
	var c = mqttclient.IntClientMqtt()
	mqttclient.StartCLientMqtt(c)
	time.Sleep(20 * time.Second)
	t.Fatal("End MQTT")
}
