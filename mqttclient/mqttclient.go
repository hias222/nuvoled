package mqttclient

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"swimdata.de/nuvoled/logging"
	"swimdata.de/nuvoled/mqttlogic"
)

var mqttServer = "localhost"
var mqttTopic = "rawdata"
var logger = logging.GetLogger()

// define a function for the default message handler
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	logger.Debug("TOPIC: " + msg.Topic())
	logger.Debug("MSG: " + string(msg.Payload()))
}

var getData MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	logger.Debug("TOPIC: " + msg.Topic())
	logger.Debug("MSG: " + string(msg.Payload()))
	data := []byte(msg.Payload())

	mqttlogic.SendUDPMessage(data)

	//mqtttoudpclient.SendUDPMessage(data)
}

func IntClientMqtt(mqttserver string) (MQTT.Client, error) {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	if mqttserver != "" {
		mqttServer = mqttserver
	}

	logger.Info("connect to tcp://" + mqttServer + ":1883")
	logger.Info("please check if it is tcp4")
	opts := MQTT.NewClientOptions().AddBroker("tcp://" + mqttServer + ":1883")

	clientid := strconv.Itoa(rand.Int())
	logger.Info("client " + clientid)

	opts.SetClientID(clientid)
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

func SubscribeMqtt(client MQTT.Client) {
	var c = client

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := c.Subscribe(mqttTopic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

}

func SendMessageMqtt(client MQTT.Client, message string) {

	var c = client

	/*
		//subscribe to the topic /go-mqtt/sample and request messages to be delivered
		//at a maximum qos of zero, wait for the receipt to confirm the subscription
		if token := c.Subscribe(mqttTopic, 0, nil); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	*/

	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
	//from the server after sending each message

	text := message
	token := c.Publish(mqttTopic, 0, false, text)
	token.Wait()

	//unsubscribe from /go-mqtt/sample

	//c.Disconnect(250)
}

func DisconnetMqtt(client MQTT.Client) {

	var c = client

	if token := c.Unsubscribe(mqttTopic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
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
