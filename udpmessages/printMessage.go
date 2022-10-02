package udpmessages

import "fmt"

var framerow int
var printDetailedDate bool
var printAllData bool

func SetParameter(printData bool, details bool) {
	printDetailedDate = printData
	printAllData = details
}

func BufferToString(data []byte, length int) {

	bufferlength := 0

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
		fmt.Print(" sync cnt ")
		fmt.Print(data[3])
		fmt.Print(" type ")
		fmt.Print(data[4])
		if data[4] == 20 {
			fmt.Print(" jpg ")
		}
		if data[4] == 10 {
			fmt.Print(" RGB888")
		}
		if data[4] == 30 {
			fmt.Print(" RGB565")
		}
		fmt.Print(" packH")
		fmt.Print(data[5])
		fmt.Print(" packL ")
		fmt.Print(data[6])
		fmt.Print(" totalH ")
		fmt.Print(data[7])
		fmt.Print(" totalL ")
		fmt.Print(data[8])
		fmt.Println("")
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

	if printDetailedDate {
		fmt.Print(framerow, ": ", bufferlength, ": ")

		if printAllData {
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
		}
		fmt.Print("\n")
	}

}
