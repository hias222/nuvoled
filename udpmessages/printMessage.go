package udpmessages

import "fmt"

func BufferToString(data []byte) {

	bufferlength := len(data)

	for i := 0; i < bufferlength; i++ {
		if i < 16 {
			fmt.Print(data[i], " ")
			//fmt.Print(string(buffer[i]), " ")
		}

	}

	fmt.Print("\n")

}
