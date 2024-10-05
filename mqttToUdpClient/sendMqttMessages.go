package mqtttoudpclient

import (
	"strconv"
	"time"

	"swimdata.de/nuvoled/image"
	"swimdata.de/nuvoled/logging"
	"swimdata.de/nuvoled/udpserver"
)

var framenumber int
var logger = logging.GetLogger()

func generateFirst10Bytes(frame int, row int, rows int, nrpackages int, buffer []byte) {
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
	buffer[8] = byte(rows)
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

	// bei 128 Zeilen ist Schluss, da muss man h füllen!!

	buffer := make([]byte, 1450)

	row := 0
	linebytesnr := 10

	calc_rows := int(len(byteRGBA)/4*3/1440) + 1

	logger.Debug("length: " + strconv.Itoa(len(byteRGBA)) + " rows " + strconv.Itoa(calc_rows))

	for i := 0; i < len(byteRGBA); i = i + 4 {

		buffer[linebytesnr] = byteRGBA[i+2]
		linebytesnr++
		buffer[linebytesnr] = byteRGBA[i+1]
		linebytesnr++
		buffer[linebytesnr] = byteRGBA[i]
		linebytesnr++

		if linebytesnr > 1449 {
			linebytesnr = 10
			generateFirst10Bytes(framenumber, row, calc_rows, 45, (buffer))
			row++
			//udpmessages.BufferToString(buffer, 1500)
			//fmt.Printf("buffer: %v\n", buffer)
			udpserver.SendUDPListenMessage(buffer)
			//time.Sleep(10 * time.Millisecond)
		}
	}

	var nrpackages = linebytesnr/32 + 1
	var linebytesnrcorrect = nrpackages * 32
	logger.Debug("endbytes: " + strconv.Itoa(linebytesnr) + " need " + strconv.Itoa(linebytesnrcorrect) + " rows " + strconv.Itoa(row))

	generateFirst10Bytes(framenumber, row, calc_rows, nrpackages, (buffer))
	bufferend := make([]byte, linebytesnrcorrect)
	for i := 0; i < linebytesnr; i++ {
		bufferend[i] = buffer[i]
	}
	for i := linebytesnr; i < linebytesnrcorrect; i++ {
		bufferend[i] = 0
	}

	//udpmessages.BufferToString(bufferend, 1450)
	udpserver.SendUDPListenMessage(bufferend)
	// Ende
	time.Sleep(10 * time.Millisecond)
	udpserver.SendUDPListenMessage(generateFrameSyncMessage(framenumber - 1))
	time.Sleep(30 * time.Millisecond)
	udpserver.SendUDPListenMessage(generateFrameSyncMessage(framenumber))
	row = 0

}

func SendEventMessage(event string, heat string) {

	framenumber++
	// add frame
	if framenumber > 255 {
		framenumber = 1
	}

	var byteRGBA = image.CreateImageRGBA(event, heat)

	logger.Debug(event)

	if event == "W 0" {
		SendUDPStartMessage()
		logger.Info("Start Message")
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
