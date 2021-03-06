package udpmessages

import "fmt"

func ActivatePanles(panel []byte) []byte {

	//<-- 36 36 160 23 49 74 (aktiviert)
	//36 36 155 0 (save config)
	// --> 36 36 15 0 74 23 49 80 52 83 32 8 8 8 8

	buffer := make([]byte, 6)

	buffer[0] = 36
	buffer[1] = 36
	buffer[2] = 160
	buffer[3] = panel[5]
	buffer[4] = panel[6]
	buffer[5] = panel[4]

	fmt.Print("ActivatePanles <-- ")
	BufferToString(buffer, 6)

	return buffer

}

func SavePanles(panel []byte) []byte {

	//<-- 36 36 160 23 49 74 (aktiviert)
	//36 36 155 0 (save config)
	// --> 36 36 15 0 74 23 49 80 52 83 32 8 8 8 8

	buffer := make([]byte, 6)

	buffer[0] = 36
	buffer[1] = 36
	buffer[2] = 155
	buffer[3] = 0

	fmt.Print("SavePanles <-- ")
	BufferToString(buffer, 6)

	return buffer
}

func TurnOnPanles(panel []byte) []byte {

	//<-- 36 36 120 2 32 8 8 1 23 49 74 8 8 0 0 (Hacken)
	//  --> 36 36 15 0 74 23 49 80 52 83 32 8 8 8 8

	buffer := make([]byte, 15)

	buffer[0] = panel[0]
	buffer[1] = panel[1]
	buffer[2] = 120
	buffer[3] = 2
	buffer[4] = 32
	buffer[5] = 8
	buffer[6] = 8
	buffer[7] = 1

	buffer[8] = panel[5]
	buffer[9] = panel[6]
	buffer[10] = panel[4]

	buffer[11] = 8
	buffer[12] = 8

	buffer[13] = 0
	buffer[14] = 0

	fmt.Print("TurnOnPanles <-- ")
	BufferToString(buffer, 15)

	return buffer

}
