package udpmessages

import "fmt"

func CreateRegisterMessage(panel []byte) []byte {
	// <-- 36 36 120 2 0 8 8 1 23 49 74 8 8 0 0
	//  --> 36 36 15 0 74 23 49 80 52 83 32 8 8 8 8

	buffer := make([]byte, 15)

	buffer[0] = panel[0]
	buffer[1] = panel[1]
	buffer[2] = 120
	buffer[3] = 2
	buffer[4] = 0
	//Größe
	buffer[5] = 8
	buffer[6] = 8
	//Nummer
	buffer[7] = 1
	//Macs
	buffer[8] = panel[5]
	buffer[9] = panel[6]
	//
	buffer[10] = panel[7]
	buffer[11] = 8
	buffer[12] = 8
	buffer[13] = 0
	buffer[14] = 0

	fmt.Print("<-- ")
	BufferToString(buffer, 15)

	return buffer

}
