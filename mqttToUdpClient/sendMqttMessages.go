package mqtttoudpclient

import (
	"fmt"
	"strings"
	"time"

	"swimdata.de/nuvoled/image"
	"swimdata.de/nuvoled/udpmessages"
	"swimdata.de/nuvoled/udpserver"
)

var framenumber int

func generateFirst10Bytes(frame int, row int, nrpackages int, buffer []byte) {
	//36 36 20 2 10 0 0 0 35 45 255 0 255
	//..
	// 36 36 20 2 10 0 6 0 35 45 255 0 255
	buffer[0] = 36
	buffer[1] = 36
	buffer[2] = 20
	buffer[3] = byte(frame)
	buffer[4] = 10
	buffer[5] = 0
	buffer[6] = byte(row)
	buffer[7] = 0
	buffer[8] = 35
	buffer[9] = byte(nrpackages)
}

func generateFrameSyncMessage(frame int) []byte {
	//36 36 100 1
	buffer := make([]byte, 4)
	buffer[0] = 36
	buffer[1] = 36
	buffer[2] = 100
	buffer[3] = byte(frame)
	return buffer

}

func SendUDPData(byteRGBA []byte, framenumber int) {

	buffer := make([]byte, 1450)

	row := 0
	linebytesnr := 10

	for i := 0; i < len(byteRGBA); i = i + 4 {

		buffer[linebytesnr] = byteRGBA[i+2]
		linebytesnr++
		buffer[linebytesnr] = byteRGBA[i+1]
		linebytesnr++
		buffer[linebytesnr] = byteRGBA[i]
		linebytesnr++

		if linebytesnr > 1449 {
			linebytesnr = 10
			generateFirst10Bytes(framenumber, row, 45, (buffer))
			row++
			//udpmessages.BufferToString(buffer, 1500)
			//fmt.Printf("buffer: %v\n", buffer)
			udpserver.SendUDPListenMessage(buffer)
			//time.Sleep(10 * time.Millisecond)
		}
	}

	var nrpackages = linebytesnr/32 + 1
	var linebytesnrcorrect = nrpackages * 32
	fmt.Println("endbytes: ", linebytesnr, " need ", linebytesnrcorrect)

	generateFirst10Bytes(framenumber, row, nrpackages, (buffer))
	bufferend := make([]byte, linebytesnrcorrect)
	for i := 0; i < linebytesnr; i++ {
		bufferend[i] = buffer[i]
	}
	for i := linebytesnr; i < linebytesnrcorrect; i++ {
		bufferend[i] = 0
	}

	udpmessages.BufferToString(bufferend, 1450)
	udpserver.SendUDPListenMessage(buffer)
	// Ende
	time.Sleep(10 * time.Millisecond)
	udpserver.SendUDPListenMessage(generateFrameSyncMessage(framenumber - 1))
	time.Sleep(30 * time.Millisecond)
	udpserver.SendUDPListenMessage(generateFrameSyncMessage(framenumber))
	row = 0

}

func sendEventMessage(event string, heat string) {

	framenumber++
	// add frame
	if framenumber > 255 {
		framenumber = 1
	}

	var byteRGBA = image.CreateImageRGBA(event, heat)

	fmt.Println(event)

	if event == "W 0" {
		SendUDPStartMessage()
		fmt.Println("Start Message")
	} else {
		SendUDPData(byteRGBA, framenumber)
	}

}

func SendUDPStartMessage() {

	framenumber++
	// add frame
	if framenumber > 255 {
		framenumber = 1
	}

	time.Sleep(1000 * time.Millisecond)

	var byteRGBA = image.GetInitImageRGBA()

	SendUDPData(byteRGBA, framenumber)

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
		sendEventMessage(event, heat)
	} else {
		fmt.Println("unknown ", messagetype)
	}

}
