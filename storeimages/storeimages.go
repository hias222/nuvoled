package storeimages

import "fmt"

type packageData struct {
	framenr    byte
	packagenrL byte
	packagenrH byte
	data       []byte
}

type frameData struct {
	pack        []packageData
	framenumber byte
}

var frames []frameData
var frameNumber byte

func SaveData(data []byte, length int) {
	//get frame number

	accessdata := data

	if accessdata[2] == 100 && length == 4 {
		storeNumber := data[3]
		for i := range frames {
			if frames[i].framenumber == storeNumber {
				storeData(frames[i].pack, frames[i].framenumber)
				frames = append(frames[:i], frames[i+1:]...)
				break
			}
		}
		return
	}

	if accessdata[2] != 20 {
		return
	}

	receivepackages := accessdata
	receiverowL := accessdata[6]
	receiverowH := accessdata[5]
	receiveframe := accessdata[3]

	//fmt.Print(receiverowH)

	var packagedata = packageData{
		// nur low
		receiveframe,
		receiverowL,
		receiverowH,
		receivepackages,
	}

	found := false

	for i := range frames {
		if frames[i].framenumber == receiveframe {
			frames[i].pack = append(frames[i].pack, packagedata)
			found = true
			break
		}
	}

	if !found {
		//add
		var packagedatas = []packageData{packagedata}
		var framedata = frameData{
			packagedatas,
			receiveframe,
		}
		frames = append(frames, framedata)
	}
}

func storeData(data []packageData, frame byte) {
	fmt.Print("frame ")
	fmt.Print(frame)
	fmt.Print(" length ")
	fmt.Println(len(data))

	var framenummer byte = 0
	for i := range data {
		if framenummer != data[i].framenr {
			fmt.Println(int(data[i].framenr))
			framenummer = data[i].framenr
		}

	}

}
