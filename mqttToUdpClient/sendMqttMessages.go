package mqtttoudpclient

import (
	"fmt"

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

func SendUDPMessage(data []byte) {
	// input Message like header 001 001
	var message = string(data)
	fmt.Println(message)

	framenumber++
	// add frame
	if framenumber > 255 {
		framenumber = 1
	}

	var byteRGBA = image.CreateImageRGBA()

	buffer := make([]byte, 1450)

	row := 0
	linebytesnr := 10

	for i := 0; i < len(byteRGBA); i = i + 4 {

		buffer[linebytesnr] = byteRGBA[i]
		linebytesnr++
		buffer[linebytesnr] = byteRGBA[i+1]
		linebytesnr++
		buffer[linebytesnr] = byteRGBA[i+2]
		linebytesnr++

		if linebytesnr > 1449 {
			linebytesnr = 10
			generateFirst10Bytes(framenumber, row, 45, (buffer))
			row++
			udpmessages.BufferToString(buffer, 1500)
			//fmt.Printf("buffer: %v\n", buffer)
			udpserver.SendUDPMessage(buffer)
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
	udpserver.SendUDPMessage(buffer)
	udpmessages.BufferToString(generateFrameSyncMessage(framenumber), 10)
	udpserver.SendUDPMessage(buffer)

}
