package udpmessages

import "fmt"

var framerow int

func BufferToString(data []byte, length int) {

	bufferlength := 0
	printdata := false

	if length > len(data) {
		bufferlength = len(data)
	} else {
		bufferlength = length
	}

	printchars := bufferlength
	printendchars := bufferlength + 1

	//fmt.Println(data[2])

	if data[2] == 20 {
		fmt.Print("[20] ")
		fmt.Println(data[3])
		printchars = 16
		printendchars = bufferlength - 9
		framerow++
	}

	if data[2] == 100 {
		fmt.Print("[100] ")
		fmt.Println(data[3])
		//fmt.Println("[100]")
		framerow = 0
	}

	if data[2] == 140 {
		fmt.Print("[140] ")
		fmt.Println(data[3])
		//fmt.Println("[100]")
		framerow = 0
	}

	if printdata {
		fmt.Println(framerow, ": ", bufferlength, ": ")

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

}
