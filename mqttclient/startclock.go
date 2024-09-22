package mqttclient

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var done = make(chan struct{})

func StartClock() {

	var c, err = IntClientMqtt("localhost")

	if err != nil {
		fmt.Println(err)
		logger.Debug("Waiting 30s")
		time.Sleep(30 * time.Second)
	} else {
		time.Sleep(1 * time.Second)
		SubscribeMqtt(c)
		SendMessageMqtt(c, "header 0 0")
		time.Sleep(5 * time.Second)
		DisconnetMqtt(c)
	}

	startTicker(c)
	logger.Info("Started Clock")
	time.Sleep(70 * time.Second)
	stopTicker()
	logger.Info("Stop Clock")

}

func startTicker(c mqtt.Client) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {

		var c, err = IntClientMqtt("localhost")

		if err != nil {
			fmt.Println(err)
			logger.Error("Error Connecting")
			return
		} else {
			logger.Info("Connected")
		}

		SubscribeMqtt(c)
		start := time.Now()

		for {
			select {
			case <-done:
				DisconnetMqtt(c)
				return
			case <-ticker.C:
				timedate := int(time.Since(start).Seconds())
				stringtime := fmt.Sprintf("clock %d", timedate)
				SendMessageMqtt(c, stringtime)
				logger.Debug(stringtime)
			}
		}

	}()

	//ticker.Stop()
	//done <- struct{}{}
	//log.Println("Done")

}

func stopTicker() {
	done <- struct{}{}
}
