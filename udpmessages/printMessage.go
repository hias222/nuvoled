package udpmessages

import (
	"fmt"
	"time"
)

var framerow int
var loop int
var printDetailedData bool
var printAllData bool
var analyseData bool
var frameTime int64

func SetParameter(printData bool, details bool, analyse bool) {
	printDetailedData = printData
	printAllData = details
	analyseData = analyse
}

func BufferToString(data []byte, length int) {

	bufferlength := 0

	newFrame := false
	lengthFrame := 0

	if length > len(data) {
		bufferlength = len(data)
	} else {
		bufferlength = length
	}

	printchars := bufferlength
	printendchars := bufferlength + 1

	//fmt.Println(data[2])

	if data[2] == 20 {
		if printDetailedData {
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
			fmt.Print(" size ")
			fmt.Print(data[9])
			fmt.Println("")
		}
		printchars = 100
		printendchars = bufferlength - 9
		newFrame = false
		framerow++
	}

	if data[2] == 100 {
		if printDetailedData {
			fmt.Print("[100] ")
			fmt.Println(data[3])
			//fmt.Println("[100]")
		}
		newFrame = true
		lengthFrame = framerow
		framerow = 0
	}

	if data[2] == 140 {
		if printDetailedData {
			fmt.Print("[140] ")
			fmt.Println(data[3])
		}
		framerow = 0
	}

	if newFrame && analyseData {
		loop++
		newFrame = false
		now := time.Now()
		nsec := now.UnixNano()
		var diff = (nsec - frameTime) / 1000 / 1000
		var hz = 1000 / diff
		frameTime = nsec

		if loop > 23 {
			loop = 0
			fmt.Print("Rows UDP ")
			fmt.Print(lengthFrame)
			fmt.Print(" fps ")
			fmt.Print(hz)
			//fmt.Print(frameTime)
			fmt.Println("")
		}
	}

	if printDetailedData {
		//fmt.Print(framerow, ": ", bufferlength, ": ")

		str1 := fmt.Sprintf("%d: %d: ", framerow, bufferlength)

		if printAllData && framerow != 0 {

			for i := 10; i < bufferlength; i++ {
				if i < printchars {
					//str1 = str1 + string(int(data[i]))
					str_neu := fmt.Sprintf("%d ", data[i])
					str1 = str1 + str_neu
					//fmt.Print(int(data[i]), " ")

					//fmt.Print(string(buffer[i]), " ")
				}

				if i == printchars {
					str1 = str1 + "..."
					//fmt.Print(" ... ")
					//fmt.Print(string(buffer[i]), " ")
				}

				if i > printendchars {
					str_neu := fmt.Sprintf("%d ", data[i])
					str1 = str1 + str_neu
					//fmt.Print(data[i], " ")
					//fmt.Print(string(buffer[i]), " ")
				}
			}
			fmt.Println(str1)
		}

	}

}
