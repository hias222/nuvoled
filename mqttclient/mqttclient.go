package mqttclient

import (
	"errors"
	"fmt"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	mqtttoudpclient "swimdata.de/nuvoled/mqttToUdpClient"
)

var mqttServer = "localhost"
var mqttTopic = "rawdata"

//define a function for the default message handler
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

var getData MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	data := []byte(msg.Payload())
	mqtttoudpclient.SendUDPMessage(data)
}

func IntClientMqtt(mqttserver string) (MQTT.Client, error) {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	if mqttserver != "" {
		mqttServer = mqttserver
	}

	fmt.Println("connect to tcp://" + mqttServer + ":1883")
	fmt.Println("please check if it is tcp4")
	opts := MQTT.NewClientOptions().AddBroker("tcp://" + mqttServer + ":1883")
	opts.SetClientID("go-simple")
	opts.SetDefaultPublishHandler(f)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		//panic(token.Error())
		return nil, errors.New(token.Error().Error())
	}

	return c, nil

}

func StartCLientMqtt(client MQTT.Client) {

	var c = client

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := c.Subscribe(mqttTopic, 0, getData); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

}

func TestClientMqtt(client MQTT.Client) {

	var c = client

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := c.Subscribe(mqttTopic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
	//from the server after sending each message
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish(mqttTopic, 0, false, text)
		token.Wait()
	}

	time.Sleep(3 * time.Second)

	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe(mqttTopic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}
