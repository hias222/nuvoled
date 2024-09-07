package mqttlogic

import (
	"fmt"
	"strings"

	mqtttoudpclient "swimdata.de/nuvoled/mqttToUdpClient"
)

func SendClock() {
	fmt.Println("xlock")
}

func getMessageType(message string) string {
	strParts := strings.Split(message, " ")
	if len(strParts) > 0 {
		return strParts[0]
	}
	return ""
}

func getEvent(message string) string {
	strParts := strings.Split(message, " ")
	if len(strParts) > 1 {
		return strParts[1]
	}
	return "000"
}

func getHeat(message string) string {
	strParts := strings.Split(message, " ")
	if len(strParts) > 2 {
		return strParts[2]
	}
	return "000"
}

func SendUDPMessage(data []byte) {
	// input Message like header 001 001
	var message = string(data)
	var messagetype = getMessageType(message)
	if messagetype == "header" {
		event := "W " + getEvent(message)
		heat := "L " + getHeat(message)
		fmt.Println("--> header event with ", event, " - ", heat)
		mqtttoudpclient.SendEventMessage(event, heat)

	} else if messagetype == "clock" {
		fmt.Println("--> clock")

		SendClock()

		//sendclock.SendClock()
	} else {
		fmt.Println("unknown ", messagetype)
	}

}
