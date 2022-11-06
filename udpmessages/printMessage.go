package udpmessages

import (
	"fmt"
	"time"

	"swimdata.de/nuvoled/storeimages"
)

var framerow int
var loop int
var printDetailedData bool
var printAllData bool
var analyseData bool
var storeData bool
var frameTime int64
var frameType int

func SetParameter(printData bool, details bool, analyse bool) {
	printDetailedData = printData
	printAllData = details
	analyseData = analyse
	storeData = true
}

func ReportFrameData(lengthFrame int, frameType int) {

	loop++
	now := time.Now()
	nsec := now.UnixNano()
	var diff = (nsec - frameTime) / 1000 / 1000
	if diff > 0 {
		var hz = 1000 / diff
		frameTime = nsec

		if loop > 23 {
			loop = 0
			fmt.Print("Rows UDP ")
			fmt.Print(lengthFrame)
			fmt.Print(" fps ")
			fmt.Print(hz)
			fmt.Print(" type ")
			fmt.Print(frameType)
			fmt.Println("")
		}
	}
}

func DetailedFrameData(printchars int, printendchars int, bufferlength int, data []byte) {
	str1 := fmt.Sprintf("row: %d length: %d type: %d H: %d L: %d", framerow, bufferlength, data[4], data[5], data[6])

	if printAllData && framerow != 0 {
		str1 = str1 + "\n"

		for i := 10; i < bufferlength; i++ {
			if i < printchars {
				str_neu := fmt.Sprintf("%d ", data[i])
				str1 = str1 + str_neu
			}

			if i == printchars {
				str1 = str1 + "..."
			}

			if i > printendchars {
				str_neu := fmt.Sprintf("%d ", data[i])
				str1 = str1 + str_neu
			}
		}
	}
	fmt.Println(str1)
}

func BufferToString(data []byte, length int) {

	bufferlength := 0
	newFrame := false
	lengthFrame := 0

	if storeData {
		storeimages.SaveData(data, length)
	}

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

		result := 0
		result += int(data[4])
		frameType = result

		printchars = 20
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
		ReportFrameData(lengthFrame, frameType)
		newFrame = false
	}

	if printDetailedData {
		DetailedFrameData(printchars, printendchars, bufferlength, data)
	}

}
