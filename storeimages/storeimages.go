package storeimages

import "fmt"

var frameNumber byte

func SaveData(data []byte, length int) {
	//get frame number

	accessdata := data

	if accessdata[2] == 100 && length == 4 {
		// save Image
		storeNumber := data[3]
		fmt.Print("Store ")
		fmt.Print(length)
		fmt.Print(" Frame ")
		fmt.Println(storeNumber)
		return
	}

	if accessdata[2] != 20 {
		return
	}

	if frameNumber != accessdata[3] {
		frameNumber = accessdata[3]
		fmt.Print("Frame ")
		fmt.Println(frameNumber)
	}

}
