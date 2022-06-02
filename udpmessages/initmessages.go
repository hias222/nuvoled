package udpmessages

import "fmt"

func CreateInitMessage() []byte {
	data := make([]byte, 15)

	// 2 mal $$
	data[0] = 36
	data[1] = 36
	// fester Wert
	data[2] = 15
	// 4 mal MAC
	data[3] = 0
	data[4] = 74
	data[5] = 23
	data[6] = 49
	// Typ 4 chars P4S
	data[7] = 80
	data[8] = 52
	data[9] = 83
	data[10] = 32
	data[11] = 8
	data[12] = 8
	data[13] = 8
	data[14] = 8

	fmt.Println("send ", data)

	return data
}
