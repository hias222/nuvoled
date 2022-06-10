package udpmessages

import "fmt"

func ResetPanles() []byte {

	buffer := make([]byte, 6)

	buffer[0] = 36
	buffer[1] = 36
	buffer[2] = 160
	buffer[3] = 0
	buffer[4] = 0
	buffer[5] = 0

	fmt.Print("reset <-- ")
	BufferToString(buffer, 6)

	return buffer

}

func RefreshPanles() []byte {

	buffer := make([]byte, 4)

	buffer[0] = 36
	buffer[1] = 36
	buffer[2] = 130
	buffer[3] = 0

	fmt.Print("refresh <-- ")
	BufferToString(buffer, 4)

	return buffer

}
