package mqttlogic

import (
	"strconv"
	"strings"

	"swimdata.de/nuvoled/logging"
	mqtttoudpclient "swimdata.de/nuvoled/mqttToUdpClient"
	"swimdata.de/nuvoled/sendclock"
)

var logger = logging.GetLogger()

func sendClock(second int) {
	sendclock.SendClock(second)
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

func getSecond(message string) int {
	strParts := strings.Split(message, " ")
	if len(strParts) > 1 {
		i, err := strconv.Atoi(strParts[1])
		if err != nil {
			return 0
		}
		return i
	}
	return 0
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
		logger.Info("--> header event with " + event + " - " + heat)
		mqtttoudpclient.SendEventMessage(event, heat)

	} else if messagetype == "clock" {
		i := getSecond(message)
		logger.Debug("--> clock " + string(i))

		sendClock(i)

		//sendclock.SendClock()
	} else {
		logger.Info("unknown ", messagetype)
	}

}
