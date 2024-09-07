package mqttclient_test

import (
	"fmt"
	"testing"
	"time"

	"swimdata.de/nuvoled/mqttclient"
)

func TestMqttClient(t *testing.T) {

	var c, err = mqttclient.IntClientMqtt("localhost")

	if err != nil {
		fmt.Println(err)
		fmt.Println("Waiting 30s")
		time.Sleep(30 * time.Second)
	} else {
		mqttclient.TestClientMqtt(c)
	}

	t.Log("End MQTT")
}

/*
func TestClient15s(t *testing.T) {
	var c = mqttclient.IntClientMqtt()
	mqttclient.StartCLientMqtt(c)
	time.Sleep(20 * time.Second)
	t.Fatal("End MQTT")
}
*/
