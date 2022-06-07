package udpmessages

import "fmt"

var framerow int

func BufferToString(data []byte, length int) {

	bufferlength := 0

	if length > len(data) {
		bufferlength = len(data)
	} else {
		bufferlength = length
	}

	printchars := bufferlength
	printendchars := bufferlength + 1

	if data[2] == 20 {
		printchars = 16
		printendchars = bufferlength - 9
		framerow++
	}

	if data[2] == 100 {
		// frame
		framerow = 0
	}

	fmt.Print(framerow, ": ", bufferlength, ": ")

	for i := 0; i < bufferlength; i++ {
		if i < printchars {
			fmt.Print(data[i], " ")
			//fmt.Print(string(buffer[i]), " ")
		}

		if i == printchars {
			fmt.Print(" ... ")
			//fmt.Print(string(buffer[i]), " ")
		}

		if i > printendchars {
			fmt.Print(data[i], " ")
			//fmt.Print(string(buffer[i]), " ")
		}
	}

	fmt.Print("\n")

}
